package main_test

import (
	"fmt"
	"strconv"

	. "github.com/chauhoangminhnguyen/material-zui-go/array"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Array", func() {
	Context("Map", func() {
		It("default", func() {
			i1 := []int{1, 2, 3}
			v1 := Map(i1, func(i int) int {
				return i + 1
			})
			Expect(v1).To(Equal([]int{2, 3, 4}))
		})
	})

	Context("ZuiArray", func() {
		It("Append", func() {
			items := &ZArray[int]{1, 2, 3}
			items2 := items.Append(4, 5, 6).Append(7)
			Expect(items2.Get()).To(Equal([]int{1, 2, 3, 4, 5, 6, 7}))
		})

		It("At", func() {
			items := ZuiArray([]int{1, 2, 3})
			Expect(items.At(1)).To(Equal(2))
			Expect(items.At(-1)).To(Equal(3))

			value, err := items.At(4)
			Expect(value).To(Equal(0))
			Expect(err).To(Equal(fmt.Errorf("Index out of range")))
		})

		It("AtByDefault", func() {
			items := ZuiArray([]int{1, 2, 3})
			Expect(items.AtByDefault(1)).To(Equal(2))
			Expect(items.AtByDefault(-1)).To(Equal(3))
			Expect(items.AtByDefault(4)).To(Equal(0))
		})

		It("AtWithDefault", func() {
			items := ZuiArray([]int{1, 2, 3})
			Expect(items.AtWithDefault(1, -1)).To(Equal(2))
			Expect(items.AtWithDefault(-1, -1)).To(Equal(3))
			Expect(items.AtWithDefault(4, -1)).To(Equal(-1))
		})

		It("Map", func() {
			items := ZArray[int]{1, 2, 3}
			mapFunc := func(i int) int { return i * 2 }
			values := items.Map(mapFunc).Map(mapFunc).Get()
			Expect(values).To(Equal([]int{4, 8, 12}))
		})

		It("MapToZuiArray", func() {
			items := []int{1, 2, 3}
			mapFunc := func(i int) float64 { return float64(i * 4) }
			values := MapToZuiArray(items, mapFunc)
			values2 := MapToZuiArray(*values, func(i float64) string { return strconv.FormatFloat(i, 'f', 1, 64) })
			Expect(values.Get()).To(Equal([]float64{4, 8, 12}))
			Expect(values2.Get()).To(Equal([]string{"4.0", "8.0", "12.0"}))
		})
	})

	Context("ZuiArrayNumber", func() {
		It("Min", func() {
			items := ZuiArrayNumber([]int{1, 2, 3})
			Expect(items.Min()).To(Equal(1))
		})

		It("Max", func() {
			float32Items := []float32{1.2, 1.21, 1.212}
			items := ZuiArrayNumber(float32Items)
			Expect(items.Max()).To(Equal(float32(1.212))) // default is float64

			items2 := ZuiArrayNumber([]float64{1.2, 1.21, 1.212})
			Expect(items2.Max()).To(Equal(1.212))
		})

		It("Sum", func() {
			items := ZuiArrayNumber([]int{1, 2, 3})
			Expect(items.Sum()).To(Equal(6))
		})

		It("Avg", func() {
			items := ZuiArrayNumber([]int{1, 2, 3})
			Expect(items.Avg()).To(Equal(2))
		})

		It("SortAsc", func() {
			items := ZuiArrayNumber([]int{3, 2, 4, 1})
			Expect(items.SortAsc().Get()).To(Equal([]int{1, 2, 3, 4}))
		})

		It("SortAsc", func() {
			items := ZuiArrayNumber([]int{3, 2, 4, 1})
			Expect(items.SortDesc().Get()).To(Equal([]int{4, 3, 2, 1}))
		})

		It("AtByDefault", func() {
			items := ZuiArrayNumber([]float32{3.1232, 2.7385, 4.324, 1.7365})
			Expect(items.AtByDefault(1).Round(2).Get()).To(Equal(2.74))
			Expect(items.AtByDefault(4).Round(2).Get()).To(Equal(float64(0)))
		})

		It("AtWithDefault", func() {
			items := ZuiArrayNumber([]float32{3.1232, 2.7385, 4.324, 1.7365})
			Expect(items.AtWithDefault(1, 0).Round(2).Get()).To(Equal(2.74))
			Expect(items.AtWithDefault(4, 1).Round(2).Get()).To(Equal(float64(1)))
		})
	})
})
