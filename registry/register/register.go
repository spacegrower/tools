package register

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/spacegrower/tools/registry/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/status"

	"github.com/spacegrower/watermelon/infra/definition"
	"github.com/spacegrower/watermelon/infra/graceful"
	"github.com/spacegrower/watermelon/infra/register"
	"github.com/spacegrower/watermelon/infra/register/etcd"
	"github.com/spacegrower/watermelon/infra/utils"
	"github.com/spacegrower/watermelon/infra/wlog"
	"github.com/spacegrower/watermelon/pkg/safe"
)

type remoteRegistry struct {
	once       sync.Once
	ctx        context.Context
	cancelFunc context.CancelFunc
	client     pb.RegistryClient
	cc         *grpc.ClientConn
	metas      []etcd.NodeMeta
	log        wlog.Logger
	reConnect  func() error
}

func NewRemoteRegister(endpoint string, opts ...grpc.DialOption) (register.ServiceRegister[etcd.NodeMeta], error) {
	ctx, cancel := context.WithCancel(context.Background())

	rr := &remoteRegistry{
		ctx:        ctx,
		cancelFunc: cancel,
		log:        wlog.With(zap.String("component", "registration")),
	}

	rr.reConnect = func() error {
		if rr.cc != nil && rr.cc.GetState() != connectivity.Ready {
			rr.cc.Close()
		}
		var err error
		rr.cc, err = grpc.DialContext(ctx, endpoint, opts...)
		if err != nil {
			return err
		}
		rr.client = pb.NewRegistryClient(rr.cc)
		return nil
	}

	if err := rr.reConnect(); err != nil {
		cancel()
		return nil, err
	}

	return rr, nil
}

func isDenied(err error) bool {
	gerr, ok := status.FromError(err)
	if !ok {
		return false
	}
	switch gerr.Code() {
	case codes.Unauthenticated:
	case codes.Aborted:
	case codes.PermissionDenied:
	default:
		return false
	}
	return true
}

func (s *remoteRegistry) Append(meta etcd.NodeMeta) error {
	// customize your register logic
	meta.Weight = utils.GetEnvWithDefault(definition.NodeWeightENVKey, 100, func(val string) (int32, error) {
		res, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
		return int32(res), nil
	})

	s.metas = append(s.metas, meta)
	return nil
}

func (s *remoteRegistry) Register() error {
	s.log.Debug("start register")

	if err := s.register(); err != nil {
		s.log.Error("failed to register service", zap.Error(err))
		if isDenied(err) {
			s.log.Error("register request is denied, exist")
			return err
		}
		if gerr, ok := status.FromError(err); ok && gerr.Code() == codes.Canceled {
			if err = s.reConnect(); err != nil {
				s.log.Error("failed to reconnect registry", zap.Error(err))
			}
		}
		return err
	}

	s.once.Do(func() {
		graceful.RegisterShutDownHandlers(func() {
			s.Close()
		})
	})

	return nil
}

func (s *remoteRegistry) parserServices() (services []*pb.ServiceInfo) {
	for _, item := range s.metas {
		meta := &pb.ServiceInfo{
			Region:      item.Region,
			OrgID:       item.OrgID,
			Namespace:   item.Namespace,
			ServiceName: item.ServiceName,
			Host:        item.Host,
			Port:        int32(item.Port),
			Weight:      item.Weight,
			Runtime:     item.Runtime,
			Tags:        item.Tags,
			Version:     item.Version,
		}

		for _, v := range item.GrpcMethods {
			meta.Methods = append(meta.Methods, &pb.MethodInfo{
				Name:           v.Name,
				IsClientStream: v.IsClientStream,
				IsServerStream: v.IsServerStream,
			})
		}

		services = append(services, meta)
	}
	return
}

func (s *remoteRegistry) register() error {
	var (
		receive  func() (*pb.Command, error)
		close    func() error
		services = s.parserServices()

		ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	)

	defer cancel()

	if len(services) == 0 {
		return errors.New("empty service")
	} else if len(services) == 1 {
		cli, err := s.client.Register(ctx)
		if err != nil {
			return err
		}

		if err = cli.Send(services[0]); err != nil {
			return err
		}

		receive = cli.Recv
		close = cli.CloseSend
	} else {
		cli, err := s.client.RegisterMulti(ctx)
		if err != nil {
			return err
		}

		if err = cli.Send(&pb.MoultiService{Services: services}); err != nil {
			return err
		}

		receive = cli.Recv
		close = cli.CloseSend
	}

	// print log
	for _, service := range services {
		s.log.Info("service registered successful",
			zap.String("namespace", service.Namespace),
			zap.String("name", service.ServiceName),
			zap.String("address", fmt.Sprintf("%s:%d", service.Host, service.Port)))
	}

	go safe.Run(func() {
		defer close()
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				resp, err := receive()
				if err != nil {
					s.log.Error("recv with error", zap.Error(err))
					if isDenied(err) {
						s.log.Error("register request is denied, exist")
						return
					}
					if gerr, ok := status.FromError(err); ok && gerr.Code() == codes.Canceled {
						time.Sleep(time.Second)
						if err = s.reConnect(); err != nil {
							s.log.Error("failed to reconnect registry", zap.Error(err))
							continue
						}
					}

					s.reRegister()
					return
				}
				s.log.Debug("TODO", zap.String("command", resp.Command), zap.Any("args", resp.Args))
			}
		}
	})

	return nil
}

func (s *remoteRegistry) DeRegister() error {
	s.cancelFunc()
	return nil
}

func (s *remoteRegistry) Close() {
	// just close kvstore not etcd client
	s.DeRegister()
	if s.cc != nil {
		s.cc.Close()
	}
}

func (s *remoteRegistry) reRegister() {
	for {
		select {
		case <-s.ctx.Done():

		default:
			if err := s.Register(); err != nil {
				time.Sleep(time.Second)
				continue
			}
		}

		return
	}
}
