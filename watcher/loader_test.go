package watcher_test

import (
	"context"
	"testing"
	"time"

	"github.com/Nomango/ark/watcher"
	"github.com/stretchr/testify/require"
)

func TestLoader(t *testing.T) {
	ch := make(chan interface{})

	alv := watcher.AutoLoad(context.Background(), watcher.NewNotifier(ch))
	require.Equal(t, nil, alv.Load())

	ch <- 1
	time.Sleep(time.Millisecond * 50)
	require.Equal(t, 1, alv.Load())

	ch <- 2
	time.Sleep(time.Millisecond * 50)
	require.Equal(t, 2, alv.Load())
}
