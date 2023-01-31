package slices

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
