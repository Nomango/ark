package ark_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/Nomango/ark"
)

func TestPtr(t *testing.T) {
	var s = "test"
	require.Equal(t, s, ark.PtrValue(ark.Ptr(s)))
	require.Equal(t, &s, ark.Ptr(s))
}
