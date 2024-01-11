package main_test

import (
	"fmt"
	"slices"

	mz_map "github.com/chauhoangminhnguyen/material-zui-go/map"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Map", func() {
	Context("Keys", func() {
		It("default", func() {
			value := map[string]int{"b": 1, "a": 2, "c": 3}
			r := mz_map.Keys(value)
			fmt.Println(value, r)
			// Expect(r).To(Equal([]string{"b", "a", "c"}))
			Expect(slices.Contains(r, "a")).To(Equal(true))
		})
	})

	Context("Values", func() {
		It("default", func() {
			value := map[string]int{"a": 1, "b": 2, "c": 3}
			r := mz_map.Values(value)
			Expect(r).To(ConsistOf(1, 2, 3))
		})
	})

	Context("Entries", func() {
		It("default", func() {
			m := map[string]string{
				"key1": "value1",
				"key2": "value2",
				"key3": "value3",
			}
			entries := mz_map.Entries(m)
			fmt.Println(m, entries)
			Expect(len(entries)).To(Equal(len(m)))

			// key1 := entries[0][0]
			// Expect(key1).To(Equal(m[key1]))

			// Expect(entries[0]).To(Equal([2]interface{}{"key1", "value1"}))
			// Expect(entries[0][0]).To(Equal("key1"))
			// Expect(entries[0][1]).To(Equal("value1"))
		})
	})

	Context("SortKey", func() {
		It("default", func() {
			i := map[string]int{"b": 2, "a": 1, "c": 3}
			r := mz_map.SortKey(i)
			Expect(r).To(Equal(map[string]int{"a": 1, "b": 2, "c": 3}))
		})
	})
})
