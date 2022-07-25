package watcher_test

import (
	"context"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Nomango/ark/watcher"
	"github.com/stretchr/testify/require"
)

func TestWatcher(t *testing.T) {
	v := int32(1)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	n := watcher.NewTimerNotifier(time.Second)
	f := func(interface{}) { atomic.AddInt32(&v, 1) }
	watcher.Watch(ctx, n, f)

	time.Sleep(time.Second / 2)
	time.Sleep(time.Second)
	require.EqualValues(t, 2, atomic.LoadInt32(&v))
	time.Sleep(time.Second)
	require.EqualValues(t, 3, atomic.LoadInt32(&v))

	cancel()
	time.Sleep(time.Second)
	require.EqualValues(t, 3, atomic.LoadInt32(&v))
}