package logs

import "context"

var DefaultLogger = defaultLogger

var defaultLogger = NewLogger()

// WithOption append options to logger
func WithOption(opts ...Option) {
	defaultLogger.WithOption(opts...)
}

// Debugf prints messages with LevelDebug
func Debugf(msg string, args ...interface{}) {
	defaultLogger.Debugf(msg, args...)
}

// Infof prints messages with LevelInfo
func Infof(msg string, args ...interface{}) {
	defaultLogger.Infof(msg, args...)
}

// Noticef prints messages with LevelNotice
func Noticef(msg string, args ...interface{}) {
	defaultLogger.Noticef(msg, args...)
}

// Warnf prints messages with LevelWarn
func Warnf(msg string, args ...interface{}) {
	defaultLogger.Warnf(msg, args...)
}

// Errorf prints messages with LevelError
func Errorf(msg string, args ...interface{}) {
	defaultLogger.Errorf(msg, args...)
}

// Debugw prints key-value pairs
func Debugw(msg string, kvs ...KeyValue) {
	defaultLogger.Debugw(msg, kvs...)
}

// Infow prints key-value pairs
func Infow(msg string, kvs ...KeyValue) {
	defaultLogger.Infow(msg, kvs...)
}

// Noticew prints key-value pairs
func Noticew(msg string, kvs ...KeyValue) {
	defaultLogger.Noticew(msg, kvs...)
}

// Warnw prints key-value pairs
func Warnw(msg string, kvs ...KeyValue) {
	defaultLogger.Warnw(msg, kvs...)
}

// Errorw prints key-value pairs
func Errorw(msg string, kvs ...KeyValue) {
	defaultLogger.Errorw(msg, kvs...)
}

// CtxDebugf prints key-value pairs first and then msgs
func CtxDebugf(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.CtxDebugf(ctx, msg, args...)
}

// CtxInfof prints key-value pairs first and then msgs
func CtxInfof(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.CtxInfof(ctx, msg, args...)
}

// CtxNoticef prints key-value pairs first and then msgs
func CtxNoticef(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.CtxNoticef(ctx, msg, args...)
}

// CtxWarnf prints key-value pairs first and then msgs
func CtxWarnf(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.CtxWarnf(ctx, msg, args...)
}

// CtxErrorf prints key-value pairs first and then msgs
func CtxErrorf(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.CtxErrorf(ctx, msg, args...)
}

// CtxDebugw prints key-value pairs
func CtxDebugw(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.CtxDebugw(ctx, msg, kvs...)
}

// CtxInfow prints key-value pairs
func CtxInfow(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.CtxInfow(ctx, msg, kvs...)
}

// CtxNoticew prints key-value pairs
func CtxNoticew(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.CtxNoticew(ctx, msg, kvs...)
}

// CtxWarnw prints key-value pairs
func CtxWarnw(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.CtxWarnw(ctx, msg, kvs...)
}

// CtxErrorw prints key-value pairs
func CtxErrorw(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.CtxErrorw(ctx, msg, kvs...)
}
