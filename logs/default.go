package logs

import "context"

var DefaultLogger = defaultLogger

var defaultLogger = NewLogger()

// WithOption append options to logger
func WithOption(opts ...Option) {
	defaultLogger.WithOption(opts...)
}

// Logf prints messages with a specific level
func Logf(level Level, msg string, args ...interface{}) {
	defaultLogger.ctxWritef(context.Background(), level, msg, args...)
}

// Debugf prints messages with LevelDebug
func Debugf(msg string, args ...interface{}) {
	defaultLogger.ctxWritef(context.Background(), LevelDebug, msg, args...)
}

// Infof prints messages with LevelInfo
func Infof(msg string, args ...interface{}) {
	defaultLogger.ctxWritef(context.Background(), LevelInfo, msg, args...)
}

// Noticef prints messages with LevelNotice
func Noticef(msg string, args ...interface{}) {
	defaultLogger.ctxWritef(context.Background(), LevelNotice, msg, args...)
}

// Warnf prints messages with LevelWarn
func Warnf(msg string, args ...interface{}) {
	defaultLogger.ctxWritef(context.Background(), LevelWarn, msg, args...)
}

// Errorf prints messages with LevelError
func Errorf(msg string, args ...interface{}) {
	defaultLogger.ctxWritef(context.Background(), LevelError, msg, args...)
}

// Logw prints key-value pairs
func Logw(level Level, msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(context.Background(), level, msg, kvs...)
}

// Debugw prints key-value pairs
func Debugw(msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(context.Background(), LevelDebug, msg, kvs...)
}

// Infow prints key-value pairs
func Infow(msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(context.Background(), LevelInfo, msg, kvs...)
}

// Noticew prints key-value pairs
func Noticew(msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(context.Background(), LevelNotice, msg, kvs...)
}

// Warnw prints key-value pairs
func Warnw(msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(context.Background(), LevelWarn, msg, kvs...)
}

// Errorw prints key-value pairs
func Errorw(msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(context.Background(), LevelError, msg, kvs...)
}

// CtxLogf prints messages with a specific level
func CtxLogf(ctx context.Context, level Level, msg string, args ...interface{}) {
	defaultLogger.ctxWritef(ctx, level, msg, args...)
}

// CtxDebugf prints key-value pairs first and then msgs
func CtxDebugf(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.ctxWritef(ctx, LevelDebug, msg, args...)
}

// CtxInfof prints key-value pairs first and then msgs
func CtxInfof(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.ctxWritef(ctx, LevelInfo, msg, args...)
}

// CtxNoticef prints key-value pairs first and then msgs
func CtxNoticef(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.ctxWritef(ctx, LevelNotice, msg, args...)
}

// CtxWarnf prints key-value pairs first and then msgs
func CtxWarnf(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.ctxWritef(ctx, LevelWarn, msg, args...)
}

// CtxErrorf prints key-value pairs first and then msgs
func CtxErrorf(ctx context.Context, msg string, args ...interface{}) {
	defaultLogger.ctxWritef(ctx, LevelError, msg, args...)
}

// CtxLogw prints key-value pairs
func CtxLogw(ctx context.Context, level Level, msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(ctx, level, msg, kvs...)
}

// CtxDebugw prints key-value pairs
func CtxDebugw(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(ctx, LevelDebug, msg, kvs...)
}

// CtxInfow prints key-value pairs
func CtxInfow(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(ctx, LevelInfo, msg, kvs...)
}

// CtxNoticew prints key-value pairs
func CtxNoticew(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(ctx, LevelNotice, msg, kvs...)
}

// CtxWarnw prints key-value pairs
func CtxWarnw(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(ctx, LevelWarn, msg, kvs...)
}

// CtxErrorw prints key-value pairs
func CtxErrorw(ctx context.Context, msg string, kvs ...KeyValue) {
	defaultLogger.ctxWritew(ctx, LevelError, msg, kvs...)
}
