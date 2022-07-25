package logs

import (
	"context"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// Valuer returns a value
type Valuer func(ctx context.Context) interface{}

// KVCallDepth returns a Valuer that returns a pkg/file:line description of the caller.
func KVCallDepth(key string, depth int) KeyValue {
	return KeyValue{Key: key, Value: valuerCallDepth(depth)}
}

// KVTimestamp returns a timestamp Valuer with a custom time format.
func KVTimestamp(key string, layout string) KeyValue {
	return KeyValue{Key: key, Value: valuerTimestamp(layout)}
}

func valuerCallDepth(depth int) Valuer {
	return func(context.Context) interface{} {
		_, file, line, _ := runtime.Caller(depth)
		idx := strings.LastIndexByte(file, '/')
		return file[idx+1:] + ":" + strconv.Itoa(line)
	}
}

func valuerTimestamp(layout string) Valuer {
	return func(context.Context) interface{} {
		return time.Now().Format(layout)
	}
}
