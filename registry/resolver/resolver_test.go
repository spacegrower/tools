package resolver

import (
	"context"
	"testing"
	"time"
)

func TestDelayer(t *testing.T) {

	type delayer struct {
		ctx    context.Context
		cancel context.CancelFunc
	}

	var rr *delayer
	if rr == nil || rr.ctx.Err() != nil {
		t.Log("o?")
	}

	rr = &delayer{}
	rr.ctx, rr.cancel = context.WithCancel(context.Background())

	go func() {
		ticker := time.NewTimer(time.Minute * 5)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				t.Log("ticker")

			case <-rr.ctx.Done():
			}
			t.Log("go return")
			return
		}
	}()

	time.Sleep(time.Second)

	rr.cancel()
	rr = nil

	time.Sleep(time.Second)

	t.Log("success")
}
