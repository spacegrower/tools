package register

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"

	"github.com/spacegrower/tools/registry/pb/space"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/spacegrower/watermelon/infra/definition"
	"github.com/spacegrower/watermelon/infra/graceful"
	"github.com/spacegrower/watermelon/infra/register"
	"github.com/spacegrower/watermelon/infra/utils"
	"github.com/spacegrower/watermelon/infra/wlog"
	"github.com/spacegrower/watermelon/pkg/safe"
)

type remoteRegistry struct {
	once       sync.Once
	ctx        context.Context
	cancelFunc context.CancelFunc
	client     space.Registry_RegisterClient
	meta       register.NodeMeta
	log        wlog.Logger
	reConnect  func() error
}

func NewRemoteRegister(endpoint string, opts ...grpc.DialOption) (register.ServiceRegister, error) {
	ctx, cancel := context.WithCancel(context.Background())

	rr := &remoteRegistry{
		ctx:        ctx,
		cancelFunc: cancel,
		log:        wlog.With(zap.String("component", "remote-register")),
	}

	rr.reConnect = func() error {
		cc, err := grpc.DialContext(ctx, endpoint, opts...)
		if err != nil {
			return err
		}
		client, err := space.NewRegistryClient(cc).Register(ctx)
		if err != nil {
			return err
		}
		rr.client = client
		return nil
	}

	if err := rr.reConnect(); err != nil {
		cancel()
		return nil, err
	}

	go safe.Run(func() {
		for {
			select {
			case <-rr.ctx.Done():
				if rr.client != nil {
					rr.client.CloseSend()
				}
				return
			default:
				resp, err := rr.client.Recv()
				if err != nil {
					rr.log.Warn("recv with error", zap.Error(err))
					if err == io.EOF {
						time.Sleep(time.Second)
						if err = rr.reConnect(); err != nil {
							rr.log.Error("failed to reconnect registry", zap.Error(err))
							continue
						}
					}

					rr.reRegister()
				} else {
					rr.log.Debug("TODO", zap.String("command", resp.Command), zap.Any("args", resp.Args))
				}
			}
		}
	})

	return rr, nil
}

func (s *remoteRegistry) Init(meta register.NodeMeta) error {
	// customize your register logic
	meta.Weight = utils.GetEnvWithDefault(definition.NodeWeightENVKey, 100, func(val string) (int32, error) {
		res, err := strconv.Atoi(val)
		if err != nil {
			return 0, err
		}
		return int32(res), nil
	})

	s.meta = meta
	return nil
}

func (s *remoteRegistry) Register() error {
	s.log.Debug("start register")

	var err error
	if err = s.register(); err != nil {
		s.log.Error("failed to register server", zap.Error(err))
		return err
	}

	s.once.Do(func() {
		graceful.RegisterShutDownHandlers(func() {
			s.Close()
		})
	})

	return nil
}

func (s *remoteRegistry) register() error {
	meta := &space.ServiceInfo{
		Region:      s.meta.Region,
		OrgID:       s.meta.OrgID,
		Namespace:   s.meta.Namespace,
		ServiceName: s.meta.ServiceName,
		Host:        s.meta.Host,
		Port:        int32(s.meta.Port),
		Weight:      s.meta.Weight,
		Runtime:     s.meta.Runtime,
		Tags:        s.meta.Tags,
		Version:     s.meta.Version,
	}

	for _, v := range s.meta.Methods {
		meta.Methods = append(meta.Methods, &space.MethodInfo{
			Name:           v.Name,
			IsClientStream: v.IsClientStream,
			IsServerStream: v.IsServerStream,
		})
	}

	if err := s.client.Send(meta); err != nil {
		if err == io.EOF {
			if err = s.reConnect(); err != nil {
				return err
			}
			return errors.New("error for retry")
		}
		return err
	}

	s.log.Info("service registered successful",
		zap.String("namespace", s.meta.Namespace),
		zap.String("name", s.meta.ServiceName),
		zap.String("address", fmt.Sprintf("%s:%d", s.meta.Host, s.meta.Port)))

	return nil
}

func (s *remoteRegistry) DeRegister() error {
	s.cancelFunc()
	return nil
}

func (s *remoteRegistry) Close() {
	// just close kvstore not etcd client
	s.DeRegister()
}

func (s *remoteRegistry) reRegister() {
	for {
		select {
		case <-s.ctx.Done():
			s.client.CloseSend()
		default:
			if err := s.Register(); err != nil {
				time.Sleep(time.Second)
				continue
			}
		}

		return
	}
}