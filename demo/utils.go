package main

func filterF[T any](test func(T) bool, in []T) (out []T) {
	for _, s := range in {
		if test(s) {
			out = append(out, s)
		}
	}
	return
}

func mapF[T, S any](f func(T) S, in []T) []S {
	out := make([]S, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}
