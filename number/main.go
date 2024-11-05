package mz_number

// func MinWithCondition[T any](s []T, compare func(T, T) int) T {
// 	r := s[0]
// 	for _, v := range s[1:] {
// 		if compare(r, v) > 0 {
// 			r = v
// 		}
// 	}
// 	return r
// }

func splitNumber(num float64) (intPart int64, fracPart float64) {
	intPart = int64(num)
	fracPart = num - float64(intPart)
	return
}
