package MR

func Map[T, E any](in []T, f func(T) E) []E {
	out := make([]E, len(in))

	for i, v := range in {
		out[i] = f(v)
	}

	return out
}

func Reduce[T, E any](in []T, start E, f func(E, T) E) E {
	out := start

	for _, v := range in {
		out = f(out, v)
	}

	return out
}

func Filter[T any](in []T, f func(T) bool) []T {
	out := make([]T, 0, len(in)) //alloc space but make empty

	for _, v := range in {
		if f(v) {
			out = append(out, v)
		}
	}

	return out
}

func AltFilter[T any](in []T, f func(T) bool) []T {
	out := make([]T, len(in))
	idx := 0

	for _, v := range in {
		if f(v) {
			out[idx] = v
			idx++
		}
	}
	return out[:idx]
}
