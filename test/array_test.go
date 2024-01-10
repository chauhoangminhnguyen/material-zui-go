package main_test

import (
	array "github.com/chauhoangminhnguyen/material-zui-go/array"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array", func() {
	Context("Map", func() {
		It("default", func() {
			i1 := []int{1, 2, 3}
			v1 := array.Map(i1, func(i int) int {
				return i + 1
			})
			Expect(v1).To(Equal([]int{2, 3, 4}))
		})
	})
})
