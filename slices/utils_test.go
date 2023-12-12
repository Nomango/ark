package slices_test

import (
	"fmt"
	"testing"

	"github.com/Nomango/ark/slices"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	result := slices.Filter([]int{1, 2, 3}, func(v int) bool { return v == 2 })
	require.Equal(t, result, []int{2})

	result = slices.Filter([]int{}, func(v int) bool { return v == 2 })
	require.Equal(t, result, []int{})

	result = slices.Filter(nil, func(v int) bool { return v == 2 })
	require.Equal(t, result, []int(nil))
}

func TestHead(t *testing.T) {
	result := slices.Head([]int{1, 2, 3}, 1)
	require.Equal(t, result, []int{1})

	result = slices.Head([]int{1, 2, 3}, 2)
	require.Equal(t, result, []int{1, 2})

	result = slices.Head([]int{1, 2, 3}, 3)
	require.Equal(t, result, []int{1, 2, 3})

	result = slices.Head([]int{1, 2, 3}, 4)
	require.Equal(t, result, []int{1, 2, 3})

	result = slices.Head([]int{1, 2, 3}, 0)
	require.Equal(t, result, []int{})
}

func TestShuffle(t *testing.T) {
	a := []int{1, 2, 4, 5, 90, 11}
	slices.Shuffle(a)
	fmt.Println("shuffle: ", a)
	slices.Shuffle(a)
	fmt.Println("shuffle again: ", a)

	b := []struct {
		Name  string
		Value int
	}{
		{"a", 1},
		{"b", 2},
		{"c", 3},
		{"d", 4},
	}

	slices.Shuffle(b)
	fmt.Println("shuffle: ", b)
	slices.Shuffle(b)
	fmt.Println("shuffle again: ", b)
}
