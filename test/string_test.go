package main_test

import (
	. "github.com/chauhoangminhnguyen/material-zui-go/string"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("String", func() {
	Context("ZuiString", func() {
		It("ToLower", func() {
			Expect(ZuiString("AbcDeF").ToLower().Get()).To(Equal("abcdef"))
		})

		It("ToUpper", func() {
			var s1 ZString = "AbcDeF"
			s2 := s1.ToLower().ToUpper()
			Expect(s2.Get()).To(Equal("ABCDEF"))
		})

		It("ReplaceWithRegex", func() {
			value := ZuiString("A1S2DF345")
			Expect(value.ReplaceWithRegex("\\D+", "").Get()).To(Equal("12345"))
		})
	})
})
