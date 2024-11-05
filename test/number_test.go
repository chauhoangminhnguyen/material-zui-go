package main_test

import (
	. "github.com/chauhoangminhnguyen/material-zui-go/number"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Number", func() {
	Context("ZuiNumber", func() {
		It("ToString", func() {
			var value ZNumber = 1 * 10
			Expect(value.ToString()).To(Equal("10"))
			Expect(ZuiNumber(12.34).ToString()).To(Equal("12.34"))
		})
	})
})
