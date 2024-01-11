package mz_array

import (
	"fmt"
	"time"
)

func ArrayTest() {
	now := time.Now()
	fmt.Println("ArrayTest", now)
}

func ArrayTest2() {
	now := time.Now()
	fmt.Println("ArrayTest2", now)
}

// define map function like js
func Map[K, V any](s []K, transform func(K) V) []V {
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}

// func at[T any](slice []T, index int) T {
// 	if len(slice) > 0 {
// 		return slice[index]
// 	}
// 	return T()
// }

func first(slice []int, defaultValue int) int {
	if len(slice) > 0 {
		return slice[0]
	}
	return defaultValue
}
