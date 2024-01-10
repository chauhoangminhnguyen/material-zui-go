package mz_number

func Min[T SignedInteger](s ...T) T {
	r := s[0]
	for _, v := range s[1:] {
		if r > v { // type is number so can be compare
			r = v
		}
	}
	return r
}

func MinWithCondition[T any](s []T, compare func(T, T) int) T {
	r := s[0]
	for _, v := range s[1:] {
		if compare(r, v) > 0 {
			r = v
		}
	}
	return r
}

func Max[T SignedInteger](s ...T) T {
	r := s[0]
	for _, v := range s[1:] {
		if r < v {
			r = v
		}
	}
	return r
}
