package slices

// Map ...
func Map[T any, U any](s []T, f func(T) U) []U {
	if s == nil {
		return []U(nil)
	}
	if len(s) == 0 {
		return []U{}
	}
	r := make([]U, 0, len(s))
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}

// Reduce ...
func Reduce[T any, U any](s []T, f func(prev U, current T) U, init U) U {
	if len(s) == 0 {
		return init
	}
	prev := init
	for _, v := range s {
		prev = f(prev, v)
	}
	return prev
}
