package security_test

import (
	"context"
	"strconv"
	"testing"
	"time"

	"github.com/spacegrower/tools/security"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestDefaultPermissionDriver(t *testing.T) {
	d := security.NewDefaultPermissionDriver([]string{"spacegrower"})
	now := time.Now().Unix()
	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(map[string]string{
		security.SignAppidKey: "test",
		security.SignSignKey:  security.GenSign("test", "testscret", now),
		security.SignTimeKey:  strconv.FormatInt(now, 10),
	}))
	if err := d.Allow(ctx); err != nil {
		if status.Convert(err).Code() != codes.PermissionDenied {
			t.Fatalf("verify error, want PermissionDenied got %s", status.Convert(err).Code().String())
		}
	}
	t.Log("success")
}
