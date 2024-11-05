package mz_array

// func ArrayTest() {
// 	now := time.Now()
// 	fmt.Println("ArrayTest", now)
// }

// define map function like js
func Map[K, V any](s []K, transform func(K) V) []V {
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}
