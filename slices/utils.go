package slices

import "math/rand"

// Filter filter slice, keep element if filter returns true
func Filter[T any](s []T, filter func(T) bool) []T {
	if len(s) == 0 {
		return s
	}
	ss := make([]T, 0, len(s))
	for _, v := range s {
		if filter(v) {
			ss = append(ss, v)
		}
	}
	return ss
}

// Head takes the first N of the slice
func Head[T any](s []T, n int) []T {
	if len(s) < n {
		return s
	}
	return s[:n]
}

func Shuffle[T any](s []T) {
	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
}
