package logs

import (
	"context"
	"fmt"
	"strconv"
	"sync/atomic"
)

type Level int

const (
	LevelDebug Level = iota
	LevelInfo
	LevelNotice
	LevelWarn
	LevelError
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "Debug"
	case LevelInfo:
		return "Info"
	case LevelNotice:
		return "Notice"
	case LevelWarn:
		return "Warn"
	case LevelError:
		return "Error"
	default:
		return "UnknownLevel" + strconv.Itoa(int(l))
	}
}

// Provider is the lowest level implementation of log writing
type Provider interface {
	Write(level Level, msg string)
}

// NewMultiProvider creates a provider that duplicates its writes to all providers
func NewMultiProvider(s ...Provider) Provider {
	return &multiProvider{providers: s}
}

type multiProvider struct {
	providers []Provider
}

func (mp *multiProvider) Write(level Level, msg string) {
	for _, p := range mp.providers {
		p.Write(level, msg)
	}
}

// NewLogger creates a logger with provider
func NewLogger(opts ...Option) *Logger {
	l := &Logger{}
	l.opt.Store(defaultOption)
	return l.WithOption(opts...)
}

// Logger is a logger implementation
type Logger struct {
	opt atomic.Value
}

// WithOption append options to logger
func (l *Logger) WithOption(opts ...Option) *Logger {
	v, _ := l.opt.Load().(option)
	for _, opt := range opts {
		opt(&v)
	}
	l.opt.Store(v)
	return l
}

// Debugf prints messages with LevelDebug
func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.ctxWrite(context.Background(), LevelDebug, msg, args...)
}

// Infof prints messages with LevelInfo
func (l *Logger) Infof(msg string, args ...interface{}) {
	l.ctxWrite(context.Background(), LevelInfo, msg, args...)
}

// Noticef prints messages with LevelNotice
func (l *Logger) Noticef(msg string, args ...interface{}) {
	l.ctxWrite(context.Background(), LevelNotice, msg, args...)
}

// Warnf prints messages with LevelWarn
func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.ctxWrite(context.Background(), LevelWarn, msg, args...)
}

// Errorf prints messages with LevelError
func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.ctxWrite(context.Background(), LevelError, msg, args...)
}

// Debugw prints key-value pairs
func (l *Logger) Debugw(msg string, kvs ...KV) {
	l.ctxWritew(context.Background(), LevelDebug, msg, kvs...)
}

// Infow prints key-value pairs
func (l *Logger) Infow(msg string, kvs ...KV) {
	l.ctxWritew(context.Background(), LevelInfo, msg, kvs...)
}

// Noticew prints key-value pairs
func (l *Logger) Noticew(msg string, kvs ...KV) {
	l.ctxWritew(context.Background(), LevelNotice, msg, kvs...)
}

// Warnw prints key-value pairs
func (l *Logger) Warnw(msg string, kvs ...KV) {
	l.ctxWritew(context.Background(), LevelWarn, msg, kvs...)
}

// Errorw prints key-value pairs
func (l *Logger) Errorw(msg string, kvs ...KV) {
	l.ctxWritew(context.Background(), LevelError, msg, kvs...)
}

// CtxDebugf prints key-value pairs first and then msgs
func (l *Logger) CtxDebugf(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWrite(ctx, LevelDebug, msg, args...)
}

// CtxInfof prints key-value pairs first and then msgs
func (l *Logger) CtxInfof(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWrite(ctx, LevelInfo, msg, args...)
}

// CtxNoticef prints key-value pairs first and then msgs
func (l *Logger) CtxNoticef(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWrite(ctx, LevelNotice, msg, args...)
}

// CtxWarnf prints key-value pairs first and then msgs
func (l *Logger) CtxWarnf(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWrite(ctx, LevelWarn, msg, args...)
}

// CtxErrorf prints key-value pairs first and then msgs
func (l *Logger) CtxErrorf(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWrite(ctx, LevelError, msg, args...)
}

// CtxDebugw prints key-value pairs
func (l *Logger) CtxDebugw(ctx context.Context, msg string, kvs ...KV) {
	l.ctxWritew(ctx, LevelDebug, msg, kvs...)
}

// CtxInfow prints key-value pairs
func (l *Logger) CtxInfow(ctx context.Context, msg string, kvs ...KV) {
	l.ctxWritew(ctx, LevelInfo, msg, kvs...)
}

// CtxNoticew prints key-value pairs
func (l *Logger) CtxNoticew(ctx context.Context, msg string, kvs ...KV) {
	l.ctxWritew(ctx, LevelNotice, msg, kvs...)
}

// CtxWarnw prints key-value pairs
func (l *Logger) CtxWarnw(ctx context.Context, msg string, kvs ...KV) {
	l.ctxWritew(ctx, LevelWarn, msg, kvs...)
}

// CtxErrorw prints key-value pairs
func (l *Logger) CtxErrorw(ctx context.Context, msg string, kvs ...KV) {
	l.ctxWritew(ctx, LevelError, msg, kvs...)
}

func (l *Logger) getOption() option {
	opt, _ := l.opt.Load().(option)
	return opt
}

func (l *Logger) ctxWritew(ctx context.Context, level Level, msg string, kvs ...KV) {
	if level < l.getOption().level {
		return
	}
	kvs = append([]KV{{Key: "msg", Value: msg}}, kvs...)
	ctx = withKV(ctx, kvs...)
	l.ctxWrite(ctx, level, "")
}

func (l *Logger) ctxWrite(ctx context.Context, level Level, msg string, args ...interface{}) {
	opt := l.getOption()
	if level < opt.level {
		return
	}
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	opt.provider.Write(level, getKVString(opt.ctx)+getKVString(ctx)+msg)
}
