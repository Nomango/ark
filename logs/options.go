package logs

import (
	"context"
)

type option struct {
	ctx      context.Context
	level    Level
	provider Provider
}

var defaultOption = option{
	ctx:      context.Background(),
	level:    LevelDebug,
	provider: &stdProvider{},
}

type Option func(*option)

// WithProvider returns an option that sets a provider
func WithProvider(p Provider) Option {
	return func(opt *option) {
		opt.provider = p
	}
}

// WithLevel returns an option that filters log level
func WithLevel(level Level) Option {
	return func(opt *option) {
		opt.level = level
	}
}

// WithKVs returns an option that sets default key-value pairs
func WithKVs(kvs ...KeyValue) Option {
	return func(opt *option) {
		opt.ctx = CtxWithKVs(opt.ctx, kvs...)
	}
}
