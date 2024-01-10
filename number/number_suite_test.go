package mz_number_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	number "github.com/chauhoangminhnguyen/material-zui-go/number"
)

func TestNumber(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TestNumber Suite")
}

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
