package ark_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Nomango/ark"
)

func TestMust(t *testing.T) {
	require.NotPanics(t, func() {
		var err error
		ark.Must(err)
	})

	require.NotPanics(t, func() {
		var err error
		require.Equal(t, 1, ark.MustValue(1, err))
	})
	require.Panics(t, func() {
		err := errors.New("some error")
		ark.Must(err)
	})

	require.Panics(t, func() {
		err := errors.New("some error")
		ark.MustValue(1, err)
	})
}
