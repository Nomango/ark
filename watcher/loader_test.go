package watcher_test

import (
	"context"
	"testing"
	"time"

	"github.com/Nomango/ark/watcher"
	"github.com/stretchr/testify/require"
)

func TestLoader(t *testing.T) {
	ch := make(chan int)

	alv := watcher.AutoLoad(context.Background(), watcher.NewNotifier(ch), watcher.WithTransformer(func(ctx context.Context, i interface{}) interface{} {
		return i.(int) + 100
	}))

	require.Equal(t, nil, alv.Load())

	ch <- 1
	time.Sleep(time.Millisecond * 50)
	require.Equal(t, 101, alv.Load())

	ch <- 2
	time.Sleep(time.Millisecond * 50)
	require.Equal(t, 102, alv.Load())
}
