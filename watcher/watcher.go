package watcher

import (
	"context"
	"time"

	"github.com/Nomango/ark/logs"
)

type Notifier interface {
	Trigger() <-chan interface{}
	Cleanup()
}

func Watch(ctx context.Context, n Notifier, f func(interface{})) {
	ctx = logs.CtxWithKVs(ctx, logs.KV("from", "watcher"))
	do := func(v interface{}) {
		defer func() {
			if e := recover(); e != nil {
				logs.CtxErrorf(ctx, "PANIC occurred!!! msg=%v", e)
			}
		}()
		f(v)
	}
	go func() {
		defer n.Cleanup()
		for {
			select {
			case v, ok := <-n.Trigger():
				if !ok {
					logs.CtxNoticef(ctx, "notifier is closed")
					return
				}
				do(v)
			case <-ctx.Done():
				logs.CtxNoticef(ctx, "context is done, err=%v", ctx.Err())
				return
			}
		}
	}()
}

func NewNotifier(trigger <-chan interface{}, cleanup func()) Notifier {
	return &notifier{
		trigger: trigger,
		cleanup: cleanup,
	}
}

func NewTimerNotifier(interval time.Duration) Notifier {
	ch := make(chan interface{})
	stopCh := make(chan struct{})
	notifier := NewNotifier(ch, func() { stopCh <- struct{}{} })

	t := time.NewTicker(interval)
	go func() {
		defer t.Stop()
		for {
			select {
			case v := <-t.C:
				ch <- v
			case <-stopCh:
				return
			}
		}
	}()
	return notifier
}

type notifier struct {
	trigger <-chan interface{}
	cleanup func()
}

func (n *notifier) Trigger() <-chan interface{} {
	return n.trigger
}

func (n *notifier) Cleanup() {
	if n.cleanup != nil {
		n.cleanup()
	}
}
