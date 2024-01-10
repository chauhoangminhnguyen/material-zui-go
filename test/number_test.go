package main_test

import (
	number "github.com/chauhoangminhnguyen/material-zui-go/number"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Number", func() {
	Context("Min", func() {
		It("default", func() {
			a := number.Min(1, 2, 3)
			Expect(a).To(Equal(1))
		})

		It("spread array", func() {
			a := number.Min([]int{1, 2, 3}...)
			Expect(a).To(Equal(1))
		})
	})
})
