package mz_array

import (
	"fmt"
	"sort"
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

func Keys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func SortKey[V any](m map[string]V) map[string]V {
	result := map[string]V{}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys) // sort keys
	for _, k := range keys {
		result[k] = m[k]
	}
	return result
}

// func Sort() {
// 	m := map[string]int{"Alice": 2, "Cecil": 1, "Bob": 3}
// 	keys := make([]string, 0, len(m))
// 	for k := range m {
// 		keys = append(keys, k)
// 	}
// 	sort.Strings(keys)

// 	for _, k := range keys {
// 		fmt.Println(k, m[k])
// 	}
// }
