package ark_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/Nomango/ark"
)

func TestWithoutCancel(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctxWithoutCancel := ark.WithoutCancel(ctx)

	cancel()

	require.Nil(t, nil, ctxWithoutCancel.Done())
	require.Nil(t, nil, ctxWithoutCancel.Err())
	ddl, ok := ctxWithoutCancel.Deadline()
	require.Empty(t, ddl)
	require.False(t, ok)
}
