package cleanup

import (
	"context"
	"sync"

	"github.com/Nomango/ark/logs"
)

var (
	funcs []func(ctx context.Context)
	mu    sync.Mutex
)

func Register(f func(ctx context.Context)) {
	mu.Lock()
	defer mu.Unlock()
	funcs = append(funcs, f)
}

func Do(ctx context.Context) {
	mu.Lock()
	tmp := funcs
	funcs = nil
	mu.Unlock()

	for _, cleanup := range tmp {
		cleanup := cleanup
		func() {
			defer func() {
				if e := recover(); e != nil {
					logs.Errorf("panic occurred in cleanup, err=", e)
				}
			}()
			cleanup(ctx)
		}()
	}
}
