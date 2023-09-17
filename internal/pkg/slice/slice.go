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
