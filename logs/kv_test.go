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

	ctx = logs.CtxWithKVs(ctx, logs.FlatKVs("3", 3, "2", 4)...)
	kvs = logs.CtxGetKVs(ctx)
	require.Len(t, kvs, 3)
	require.Equal(t, kvs[0], logs.KV{"1", 1})
	require.Equal(t, kvs[1], logs.KV{"2", 4})
	require.Equal(t, kvs[2], logs.KV{"3", 3})
}
