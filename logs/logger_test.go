package logs_test

import (
	"context"
	"testing"
	"time"

	"github.com/Nomango/ark/logs"
)

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
	logs.WithOption(logs.WithKVs(logs.KVTimestamp("time", time.RFC3339), logs.KVCallDepth("caller", 4)))
	ctx = logs.CtxWithKVs(ctx, logs.KV{"test", 1})

	logs.CtxDebugf(ctx, "test msg, v=%d", 1)
	logs.CtxDebugw(ctx, "test msg,", logs.KV{"v", 1})

	logs.WithOption(logs.WithLevel(logs.LevelInfo))
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
