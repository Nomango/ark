package slices_test

import (
	"strconv"
	"testing"

	"github.com/Nomango/ark/slices"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	result := slices.Map([]int{1, 2, 3}, func(v int) int {
		return v + 1
	})
	require.Equal(t, result, []int{2, 3, 4})

	result = slices.Map([]string{"1", "2", "3"}, func(v string) int {
		n, _ := strconv.ParseInt(v, 10, 64)
		return int(n)
	})
	require.Equal(t, result, []int{1, 2, 3})

	result = slices.Map([]int{}, func(v int) int {
		return v + 1
	})
	require.Equal(t, result, []int{})

	result = slices.Map(nil, func(v int) int {
		return v + 1
	})
	require.Equal(t, result, []int(nil))
}

func TestReduce(t *testing.T) {
	result := slices.Reduce([]int{1, 2, 3}, func(prev, current int) int {
		return prev + current
	}, 10)
	require.Equal(t, result, 16)

	result = slices.Reduce([]string{"1", "2", "3"}, func(prev int, current string) int {
		n, _ := strconv.ParseInt(current, 10, 64)
		return prev + int(n)
	}, 20)
	require.Equal(t, result, 26)

	result = slices.Reduce([]int{}, func(prev, current int) int {
		return prev + current
	}, 30)
	require.Equal(t, result, 30)

	result = slices.Reduce(nil, func(prev, current int) int {
		return prev + current
	}, 40)
	require.Equal(t, result, 40)
}
