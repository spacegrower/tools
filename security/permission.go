package security

import (
	"context"

	"github.com/spacegrower/watermelon/infra/middleware"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type PermissionDriver interface {
	Allow(ctx context.Context) error
}

type DefaultPermissionDriver struct {
	allows map[string]bool
}

func (d *DefaultPermissionDriver) Allow(ctx context.Context) error {
	md, exist := metadata.FromIncomingContext(ctx)
	if !exist {
		return status.Error(codes.Unauthenticated, "metadata is not found")
	}

	info, err := ParseSignFromIncommingMetadata(md)
	if err != nil {
		return err
	}

	if !d.allows[info.Appid()] {
		return status.Error(codes.PermissionDenied, codes.PermissionDenied.String())
	}
	return nil
}

func NewDefaultPermissionDriver(allowsAppid []string) *DefaultPermissionDriver {
	p := &DefaultPermissionDriver{
		allows: make(map[string]bool),
	}
	for _, v := range allowsAppid {
		p.allows[v] = true
	}
	return p
}

func UsePermissionDriver(d PermissionDriver) middleware.Middleware {
	return func(ctx context.Context) (err error) {
		if err = d.Allow(ctx); err != nil {
			_, ok := status.FromError(err)
			if !ok {
				return status.Error(codes.Unauthenticated, err.Error())
			}
		}
		return
	}
}
