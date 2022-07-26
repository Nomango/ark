package watcher

import (
	"context"
	"sync/atomic"
)

// AutoLoad updates returned atomic value every time n triggers
func AutoLoad(ctx context.Context, n Notifier) *atomic.Value {
	l := NewLoader(n)
	l.Start(ctx)
	return &l.v
}

type Loader struct {
	w *Watcher
	v atomic.Value
}

func NewLoader(n Notifier) *Loader {
	l := &Loader{}
	l.w = NewWatcher(n, l.receive)
	return l
}

func (l *Loader) Get() interface{} {
	return l.v.Load()
}

func (l *Loader) Start(ctx context.Context) {
	l.w.Start(ctx)
}

func (l *Loader) Stop() {
	l.w.Stop()
}

func (l *Loader) receive(v interface{}) {
	l.v.Store(v)
}
