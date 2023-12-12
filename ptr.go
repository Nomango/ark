package ark

func Ptr[T any](t T) *T {
	return &t
}

func PtrValue[T any](t *T) T {
	if t == nil {
		var empty T
		return empty
	}
	return *t
}
