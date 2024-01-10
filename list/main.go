package mz_list

import (
	"fmt"
	"time"
)

func ListTest() {
	now := time.Now()
	fmt.Println("ListTest", now)
}

func ListTest2() {
	now := time.Now()
	fmt.Println("ListTest2", now)
}

// define map function like js
func Map[K, V any](s []K, transform func(K) V) []V {
	rs := make([]V, len(s))
	for i, v := range s {
		rs[i] = transform(v)
	}
	return rs
}
