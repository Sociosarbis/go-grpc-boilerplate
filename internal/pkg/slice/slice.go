package slice

func Map[T any, K any, F func(item T) K](in []T, fn F) []K {
	if in == nil {
		return nil
	}
	var s = make([]K, len(in))
	for i, t := range in {
		s[i] = fn(t)
	}

	return s
}

func ToMap[T comparable](in []T) map[T]any {
	m := make(map[T]any, len(in))
	var empty interface{}
	for _, v := range in {
		m[v] = empty
	}
	return m
}
