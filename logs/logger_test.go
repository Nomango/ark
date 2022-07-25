package logs_test

import (
	"context"
	"testing"

	"github.com/Nomango/ark/logs"
	"github.com/stretchr/testify/require"
)

func TestKV(t *testing.T) {
	ctx := context.Background()

	kvs := logs.CtxGetKVs(ctx)
	require.Empty(t, kvs)

	ctx = logs.CtxWithKVs(ctx, logs.KV{"1", 1}, logs.KV{"2", 2})
	kvs = logs.CtxGetKVs(ctx)
	require.Len(t, kvs, 2)
	require.Equal(t, kvs[0], logs.KV{"1", 1})
	require.Equal(t, kvs[1], logs.KV{"2", 2})

	ctx = logs.CtxWithKVs(ctx, logs.KV{"3", 3}, logs.KV{"2", 4})
	kvs = logs.CtxGetKVs(ctx)
	require.Len(t, kvs, 3)
	require.Equal(t, kvs[0], logs.KV{"1", 1})
	require.Equal(t, kvs[1], logs.KV{"2", 4})
	require.Equal(t, kvs[2], logs.KV{"3", 3})
}

func TestLog(t *testing.T) {
	logs.Debugf("test msg, v=%d", 1)
	logs.Debugw("test msg,", logs.KV{"v", 1})
	logs.Infof("test msg, v=%d", 1)
	logs.Infow("test msg,", logs.KV{"v", 1})
	logs.Noticef("test msg, v=%d", 1)
	logs.Noticew("test msg,", logs.KV{"v", 1})
	logs.Warnf("test msg, v=%d", 1)
	logs.Warnw("test msg,", logs.KV{"v", 1})
	logs.Errorf("test msg, v=%d", 1)
	logs.Errorw("test msg,", logs.KV{"v", 1})

	ctx := context.Background()
	ctx = logs.CtxWithKVs(ctx, logs.KV{"1", 1}, logs.KV{"2", 2})
	logs.CtxDebugf(ctx, "test msg, v=%d", 1)
	logs.CtxDebugw(ctx, "test msg,", logs.KV{"v", 1})
	logs.CtxInfof(ctx, "test msg, v=%d", 1)
	logs.CtxInfow(ctx, "test msg,", logs.KV{"v", 1})
	logs.CtxNoticef(ctx, "test msg, v=%d", 1)
	logs.CtxNoticew(ctx, "test msg,", logs.KV{"v", 1})
	logs.CtxWarnf(ctx, "test msg, v=%d", 1)
	logs.CtxWarnw(ctx, "test msg,", logs.KV{"v", 1})
	logs.CtxErrorf(ctx, "test msg, v=%d", 1)
	logs.CtxErrorw(ctx, "test msg,", logs.KV{"v", 1})
}
