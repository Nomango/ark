package logs

import (
	"context"
	"fmt"
	"strings"
)

// KeyValue is a key-value pair
type KeyValue struct {
	Key   string
	Value interface{}
}

func KV(key string, value interface{}) KeyValue {
	return KeyValue{
		Key:   key,
		Value: value,
	}
}

// CtxWithKVs append key-value pairs to context
func CtxWithKVs(ctx context.Context, kvs ...KeyValue) context.Context {
	return withKV(ctx, kvs...)
}

// CtxGetKVs returns key-value pairs from context
func CtxGetKVs(ctx context.Context) []KeyValue {
	return getKVs(ctx)
}

// FlatKVs returns key-value pairs from flat interface slice
func FlatKVs(kvis ...interface{}) []KeyValue {
	if len(kvis)%2 != 0 {
		kvis = append(kvis, "%MISSING%")
	}
	kvs := make([]KeyValue, 0, len(kvis))
	for i := 0; i < len(kvis); i += 2 {
		kvs = append(kvs, KeyValue{fmt.Sprint(kvis[i]), kvis[i+1]})
	}
	return kvs
}

type ctxKeyKV string

func withKV(ctx context.Context, kvs ...KeyValue) context.Context {
	key := ctxKeyKV("")
	if oldkvs, ok := ctx.Value(key).([]KeyValue); ok {
		newkvs := make([]KeyValue, 0, len(oldkvs)+len(kvs))
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

func getKVs(ctx context.Context) []KeyValue {
	key := ctxKeyKV("")
	kvs, _ := ctx.Value(key).([]KeyValue)
	return kvs
}

type ctxKeyKVString string

func withKVString(ctx context.Context, kvs []KeyValue) context.Context {
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
