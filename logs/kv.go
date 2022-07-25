package logs

import (
	"context"
	"fmt"
	"strings"
)

// KV is a key-value pair
type KV struct {
	Key   string
	Value interface{}
}

// CtxWithKVs append key-value pairs to context
func CtxWithKVs(ctx context.Context, kvs ...KV) context.Context {
	return withKV(ctx, kvs...)
}

// CtxGetKVs returns key-value pairs from context
func CtxGetKVs(ctx context.Context) []KV {
	return getKVs(ctx)
}

type ctxKeyKV string

func withKV(ctx context.Context, kvs ...KV) context.Context {
	key := ctxKeyKV("")
	if oldkvs, ok := ctx.Value(key).([]KV); ok {
		newkvs := make([]KV, 0, len(oldkvs)+len(kvs))
		exists := make(map[string]int)
		for idx, kv := range oldkvs {
			newkvs = append(newkvs, kv)
			exists[kv.Key] = idx
		}
		for _, kv := range kvs {
			if idx, ok := exists[kv.Key]; ok {
				newkvs[idx] = kv
				continue
			}
			exists[kv.Key] = len(newkvs)
			newkvs = append(newkvs, kv)
		}
		kvs = newkvs
	}
	ctx = withKVString(ctx, kvs)
	return context.WithValue(ctx, key, kvs)
}

func getKVs(ctx context.Context) []KV {
	key := ctxKeyKV("")
	kvs, _ := ctx.Value(key).([]KV)
	return kvs
}

type ctxKeyKVString string

func withKVString(ctx context.Context, kvs []KV) context.Context {
	key := ctxKeyKVString("")
	hasValuer := false
	for _, kv := range kvs {
		if _, ok := kv.Value.(Valuer); ok {
			hasValuer = true
			break
		}
	}
	if hasValuer {
		// without cache
		return context.WithValue(ctx, key, nil)
	}
	// with cache
	var sb strings.Builder
	for _, kv := range kvs {
		fmt.Fprintf(&sb, "%s=%v ", kv.Key, kv.Value)
	}
	return context.WithValue(ctx, key, sb.String())
}

func getKVString(ctx context.Context) string {
	if cache, ok := ctx.Value(ctxKeyKVString("")).(string); ok {
		return cache
	}
	kvs := getKVs(ctx)
	if len(kvs) == 0 {
		return ""
	}
	var sb strings.Builder
	for _, kv := range kvs {
		if valuer, ok := kv.Value.(Valuer); ok {
			fmt.Fprintf(&sb, "%s=%v ", kv.Key, valuer(ctx))
		} else {
			fmt.Fprintf(&sb, "%s=%v ", kv.Key, kv.Value)
		}
	}
	return sb.String()
}
