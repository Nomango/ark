package watcher

import (
	"context"
	"sync/atomic"
)

// AutoLoad updates returned atomic value every time n triggers
func AutoLoad(ctx context.Context, n Notifier, opts ...LoaderOption) *atomic.Value {
	l := NewLoader(n, opts...)
	l.Start(ctx)
	return &l.v
}

type Transformer = func(context.Context, interface{}) interface{}

type Loader struct {
	w *Watcher
	v atomic.Value

	transforms []Transformer
}

func NewLoader(n Notifier, opts ...LoaderOption) *Loader {
	l := &Loader{}
	l.w = NewWatcher(n, l.receive)
	for _, opt := range opts {
		opt(l)
	}
	return l
}

type LoaderOption func(*Loader)

func WithTransformer(t Transformer) LoaderOption {
	return func(l *Loader) {
		l.transforms = append(l.transforms, t)
	}
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

func (l *Loader) receive(ctx context.Context, v interface{}) {
	for _, t := range l.transforms {
		v = t(ctx, v)
	}
	l.v.Store(v)
}
