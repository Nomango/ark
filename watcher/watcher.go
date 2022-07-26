package watcher

import (
	"context"

	"github.com/Nomango/ark/logs"
)

// Watch executes f every time n triggers
func Watch(ctx context.Context, n Notifier, f func(interface{})) {
	NewWatcher(n, f).Start(ctx)
}

type Watcher struct {
	n    Notifier
	f    func(interface{})
	stop chan struct{}
}

func NewWatcher(n Notifier, f func(interface{})) *Watcher {
	return &Watcher{
		n:    n,
		f:    f,
		stop: make(chan struct{}),
	}
}

func (w *Watcher) Start(ctx context.Context) {
	ctx = logs.CtxWithKVs(ctx, logs.KV("from", "watcher"))
	do := func(v interface{}) {
		defer func() {
			if e := recover(); e != nil {
				logs.CtxErrorf(ctx, "PANIC occurred!!! msg=%v", e)
			}
		}()
		w.f(v)
	}
	go func() {
		defer w.n.Cleanup()
		for {
			select {
			case v, ok := <-w.n.Trigger():
				if !ok {
					logs.CtxNoticef(ctx, "notifier is closed")
					return
				}
				do(v)
			case <-ctx.Done():
				logs.CtxNoticef(ctx, "context is done, err=%v", ctx.Err())
				return
			case <-w.stop:
				logs.CtxNoticef(ctx, "watcher is stoped")
				return
			}
		}
	}()
}

func (w *Watcher) Stop() {
	w.stop <- struct{}{}
}
