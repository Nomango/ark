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

// Logf prints messages with a specific level
func (l *Logger) Logf(level Level, msg string, args ...interface{}) {
	l.ctxWritef(context.Background(), level, msg, args...)
}

// Debugf prints messages with LevelDebug
func (l *Logger) Debugf(msg string, args ...interface{}) {
	l.ctxWritef(context.Background(), LevelDebug, msg, args...)
}

// Infof prints messages with LevelInfo
func (l *Logger) Infof(msg string, args ...interface{}) {
	l.ctxWritef(context.Background(), LevelInfo, msg, args...)
}

// Noticef prints messages with LevelNotice
func (l *Logger) Noticef(msg string, args ...interface{}) {
	l.ctxWritef(context.Background(), LevelNotice, msg, args...)
}

// Warnf prints messages with LevelWarn
func (l *Logger) Warnf(msg string, args ...interface{}) {
	l.ctxWritef(context.Background(), LevelWarn, msg, args...)
}

// Errorf prints messages with LevelError
func (l *Logger) Errorf(msg string, args ...interface{}) {
	l.ctxWritef(context.Background(), LevelError, msg, args...)
}

// Logw prints key-value pairs
func (l *Logger) Logw(level Level, msg string, kvs ...KeyValue) {
	l.ctxWritew(context.Background(), level, msg, kvs...)
}

// Debugw prints key-value pairs
func (l *Logger) Debugw(msg string, kvs ...KeyValue) {
	l.ctxWritew(context.Background(), LevelDebug, msg, kvs...)
}

// Infow prints key-value pairs
func (l *Logger) Infow(msg string, kvs ...KeyValue) {
	l.ctxWritew(context.Background(), LevelInfo, msg, kvs...)
}

// Noticew prints key-value pairs
func (l *Logger) Noticew(msg string, kvs ...KeyValue) {
	l.ctxWritew(context.Background(), LevelNotice, msg, kvs...)
}

// Warnw prints key-value pairs
func (l *Logger) Warnw(msg string, kvs ...KeyValue) {
	l.ctxWritew(context.Background(), LevelWarn, msg, kvs...)
}

// Errorw prints key-value pairs
func (l *Logger) Errorw(msg string, kvs ...KeyValue) {
	l.ctxWritew(context.Background(), LevelError, msg, kvs...)
}

// CtxLogf prints messages with a specific level
func (l *Logger) CtxLogf(ctx context.Context, level Level, msg string, args ...interface{}) {
	l.ctxWritef(ctx, level, msg, args...)
}

// CtxDebugf prints key-value pairs first and then msgs
func (l *Logger) CtxDebugf(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWritef(ctx, LevelDebug, msg, args...)
}

// CtxInfof prints key-value pairs first and then msgs
func (l *Logger) CtxInfof(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWritef(ctx, LevelInfo, msg, args...)
}

// CtxNoticef prints key-value pairs first and then msgs
func (l *Logger) CtxNoticef(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWritef(ctx, LevelNotice, msg, args...)
}

// CtxWarnf prints key-value pairs first and then msgs
func (l *Logger) CtxWarnf(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWritef(ctx, LevelWarn, msg, args...)
}

// CtxErrorf prints key-value pairs first and then msgs
func (l *Logger) CtxErrorf(ctx context.Context, msg string, args ...interface{}) {
	l.ctxWritef(ctx, LevelError, msg, args...)
}

// CtxLogw prints key-value pairs
func (l *Logger) CtxLogw(ctx context.Context, level Level, msg string, kvs ...KeyValue) {
	l.ctxWritew(ctx, level, msg, kvs...)
}

// CtxDebugw prints key-value pairs
func (l *Logger) CtxDebugw(ctx context.Context, msg string, kvs ...KeyValue) {
	l.ctxWritew(ctx, LevelDebug, msg, kvs...)
}

// CtxInfow prints key-value pairs
func (l *Logger) CtxInfow(ctx context.Context, msg string, kvs ...KeyValue) {
	l.ctxWritew(ctx, LevelInfo, msg, kvs...)
}

// CtxNoticew prints key-value pairs
func (l *Logger) CtxNoticew(ctx context.Context, msg string, kvs ...KeyValue) {
	l.ctxWritew(ctx, LevelNotice, msg, kvs...)
}

// CtxWarnw prints key-value pairs
func (l *Logger) CtxWarnw(ctx context.Context, msg string, kvs ...KeyValue) {
	l.ctxWritew(ctx, LevelWarn, msg, kvs...)
}

// CtxErrorw prints key-value pairs
func (l *Logger) CtxErrorw(ctx context.Context, msg string, kvs ...KeyValue) {
	l.ctxWritew(ctx, LevelError, msg, kvs...)
}

func (l *Logger) getOption() option {
	opt, _ := l.opt.Load().(option)
	return opt
}

func (l *Logger) ctxWritew(ctx context.Context, level Level, msg string, kvs ...KeyValue) {
	if level < l.getOption().level {
		return
	}
	kvs = append([]KeyValue{{Key: "msg", Value: msg}}, kvs...)
	ctx = withKV(ctx, kvs...)
	l.ctxWrite(ctx, level, "")
}

func (l *Logger) ctxWritef(ctx context.Context, level Level, msg string, args ...interface{}) {
	if level < l.getOption().level {
		return
	}
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	l.ctxWrite(ctx, level, msg)
}

func (l *Logger) ctxWrite(ctx context.Context, level Level, msg string) {
	opt := l.getOption()
	if level < opt.level {
		return
	}
	opt.provider.Write(level, getKVString(opt.ctx)+getKVString(ctx)+msg)
}
