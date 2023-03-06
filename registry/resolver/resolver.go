package resolver

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"path/filepath"
	"sync"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/attributes"
	"google.golang.org/grpc/resolver"

	"github.com/spacegrower/tools/registry/pb"
	"github.com/spacegrower/watermelon/infra"
	"github.com/spacegrower/watermelon/infra/register"
	"github.com/spacegrower/watermelon/infra/register/etcd"
	wresolver "github.com/spacegrower/watermelon/infra/resolver"
	"github.com/spacegrower/watermelon/infra/wlog"
	"github.com/spacegrower/watermelon/pkg/safe"
)

const (
	RemoteResolverScheme = "watermeloneremote"
)

func NewRemoteResolver(endpoint string, opts ...grpc.DialOption) (wresolver.Resolver, error) {
	ctx, cancel := context.WithCancel(context.Background())

	cc, err := grpc.DialContext(ctx, endpoint, opts...)
	if err != nil {
		cancel()
		return nil, err
	}

	rr := &remoteRegistry{
		ctx:    ctx,
		cancel: cancel,
		log:    wlog.With(zap.String("component", "remote-resolver-builder")),
		client: pb.NewRegistryClient(cc),
	}

	resolver.Register(rr)

	return rr, nil
}

type remoteRegistry struct {
	ctx         context.Context
	cancel      context.CancelFunc
	client      pb.RegistryClient
	grpcOptions []grpc.DialOption
	namespace   string
	log         wlog.Logger
}

func (r *remoteRegistry) Scheme() string {
	return RemoteResolverScheme
}

func (r *remoteRegistry) GenerateTarget(fullServiceName string) string {
	return fmt.Sprintf("%s://%s/%s", RemoteResolverScheme, "", fullServiceName)
}

func (r *remoteRegistry) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (
	resolver.Resolver, error) {

	service := filepath.Base(filepath.ToSlash(target.URL.Path))

	ctx, cancel := context.WithCancel(r.ctx)
	rr := &remoteResolver{
		ctx:     ctx,
		cancel:  cancel,
		client:  r.client,
		service: service,
		region:  target.URL.Query().Get("region"),
		target:  target,
		cc:      cc,
		opts:    opts,
		log:     wlog.With(zap.String("component", "remote-resolver")),
	}

	wlog.Debug("will find target", zap.String("target", target.Endpoint()))

	rr.onResolve = func(resp *pb.ResolveInfo) {
		var addrs []resolver.Address
		config, err := parseServiceConfig(resp.Config)
		if err != nil {
			rr.log.Error("failed to parse service config", zap.Error(err), zap.String("service", service))
			addrs = []resolver.Address{
				wresolver.NilAddress,
			}
		} else {
			rr.log.Debug("new receive config", zap.String("config", string(resp.Config)))

			if config.Disabled {
				addrs = []resolver.Address{
					wresolver.Disabled,
				}
			} else {
				addrs, err = rr.resolve(resp)
				if err != nil {
					r.log.Error("failed to resolve service addresses", zap.Error(err), zap.String("service", service))
					addrs = []resolver.Address{
						wresolver.NilAddress,
					}
				} else {
					rr.log.Debug("new receive address", zap.Any("address", addrs))
				}
			}
		}

		if len(addrs) == 1 && addrs[0] == wresolver.NilAddress && (rr.delayer == nil || rr.delayer.ctx.Err() != nil) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Minute*3)
			rr.delayer = &delayer{
				ctx:    ctx,
				cancel: cancel,
			}
			go func() {
				ticker := time.NewTimer(time.Minute * 5)
				defer ticker.Stop()
				for {
					select {
					case <-ticker.C:
						cc.UpdateState(resolver.State{
							Addresses: []resolver.Address{wresolver.NilAddress},
						})

					case <-rr.delayer.ctx.Done():
					}
					return
				}
			}()
			return
		}

		if rr.delayer.ctx != nil {
			rr.delayer.cancel()
			rr.delayer = nil
		}

		if err := cc.UpdateState(resolver.State{
			Addresses:     addrs,
			ServiceConfig: cc.ParseServiceConfig(wresolver.ParseCustomizeToGrpcServiceConfig(config)),
		}); err != nil {
			rr.log.Error("failed to update connect state", zap.Error(err))
		}
	}

	if err := watch(rr); err != nil {
		rr.log.Error("failed to resolver", zap.Error(err))
	}

	return rr, nil
}

type remoteResolver struct {
	ctx    context.Context
	cancel context.CancelFunc

	client pb.RegistryClient

	service   string
	region    string
	target    resolver.Target
	cc        resolver.ClientConn
	opts      resolver.BuildOptions
	run       func() (revc func() (*pb.ResolveInfo, error), update func(*pb.ResolveInfo))
	log       wlog.Logger
	onResolve func(resp *pb.ResolveInfo)

	locker sync.Mutex

	delayer *delayer
}

type delayer struct {
	ctx    context.Context
	cancel func()
}

func (r *remoteResolver) ResolveNow(_ resolver.ResolveNowOptions) {
	if r.ctx == nil || r.ctx.Err() != nil {
		if err := watch(r); err != nil {
			r.log.Error("failed to resolver", zap.Error(err))
		}
		return
	}
}
func (r *remoteResolver) Close() {
	r.cancel()
}

func (r *remoteResolver) resolve(resp *pb.ResolveInfo) ([]resolver.Address, error) {
	var result []resolver.Address
	for addr, conf := range resp.Address {
		addr, err := parseNodeInfo(addr, conf, func(attr etcd.NodeMeta, addr *resolver.Address) bool {
			if r.region == "" {
				return true
			}

			if attr.Region != r.region {
				proxy := infra.ResolveProxy(attr.Region)
				if proxy == "" {
					return false
				}

				addr.Addr = proxy
			}
			return true

		})
		if err != nil {
			if err != filterError {
				r.log.Error("parse node info with error", zap.Error(err))
			}
			continue
		}
		result = append(result, addr)
	}

	if len(result) == 0 {
		return []resolver.Address{wresolver.NilAddress}, nil
	}

	return result, nil
}

func parseServiceConfig(val []byte) (*wresolver.CustomizeServiceConfig, error) {
	var result wresolver.CustomizeServiceConfig
	if err := json.Unmarshal(val, &result); err != nil {
		return nil, err
	}

	if result.Balancer == "" {
		return nil, errors.New("service config is wrong")
	}

	return &result, nil
}

var filterError = errors.New("filter")

func parseNodeInfo(key string, val []byte, allowFunc func(attr etcd.NodeMeta, addr *resolver.Address) bool) (resolver.Address, error) {
	addr := resolver.Address{Addr: filepath.ToSlash(filepath.Base(string(key)))}
	var attr etcd.NodeMeta
	if err := json.Unmarshal(val, &attr); err != nil {
		return addr, err
	}

	if ok := allowFunc(attr, &addr); !ok {
		return addr, filterError
	}

	addr.BalancerAttributes = attributes.New(register.NodeMetaKey{}, attr)

	return addr, nil
}

func watch(r *remoteResolver) error {

	r.locker.Lock()
	defer r.locker.Unlock()

	r.ctx, r.cancel = context.WithCancel(context.Background())
	remoteResolver, err := r.client.Resolver(r.ctx,
		&pb.TargetInfo{
			Service: r.target.URL.Path,
		})
	if err != nil {
		return err
	}

	go safe.Run(func() {
		for {
			select {
			case <-r.ctx.Done():
				return
			default:
			}

			recv, err := remoteResolver.Recv()
			if err != nil {
				r.log.Error("failed to resolve service address", zap.Error(err))
				// rr.onUpdate(&space.ResolveInfo{})
				r.Close()
				return
			}

			r.onResolve(recv)
		}
	})
	return nil
}
