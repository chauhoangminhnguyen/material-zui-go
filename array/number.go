package mz_array

import (
	. "github.com/chauhoangminhnguyen/material-zui-go/number"
)

type ZAllNumber interface {
	int | int8 | int16 | int32 | int64 | float32 | float64 // define generic type must be one of these type
}

// implement mixin, like inheritance
type ZArrayNumber[T ZAllNumber] struct {
	ZArray[T]
}

func ZuiArrayNumber[T ZAllNumber](items []T) *ZArrayNumber[T] {
	result := &ZArrayNumber[T]{}
	for _, value := range items {
		result.Append(value)
	}
	return result
}

func (items *ZArrayNumber[T]) GetFloat64() []float64 {
	results := make([]float64, items.Length())
	for _, v := range items.Get() {
		results = append(results, float64(v))
	}
	return results
}

/*
AtWithDefault returns the element at the specified index in the ZArrayNumber.

It takes an integer index and a default value of type T as parameters.

The function returns a pointer to a ZNumber.
*/
func (items *ZArrayNumber[T]) AtWithDefault(index int, defaultValue T) *ZNumber {
	value, err := items.At(index)
	if err != nil {
		return ZuiNumber(float64(defaultValue))
	}
	return ZuiNumber(float64(value))
}

/*
AtByDefault returns the element at the specified index in the ZArrayNumber.

It takes an integer index as a parameter and returns a pointer to a ZNumber.
*/
func (items *ZArrayNumber[T]) AtByDefault(index int) *ZNumber {
	value, _ := items.At(index)
	return ZuiNumber(float64(value))
}

func (items *ZArrayNumber[T]) Min() T {
	length := items.Length()
	if length == 0 {
		return *new(T)
	}

	min := items.First()
	if length == 1 {
		return min
	}

	values := (items.Get())[1:]
	for _, v := range values {
		if min > v { // type is number so can be compare
			min = v
		}
	}
	return min
}

func (items *ZArrayNumber[T]) Max() T {
	length := items.Length()
	if length == 0 {
		return *new(T)
	}

	max := items.First()
	if length == 1 {
		return max
	}

	values := (items.Get())[1:]
	for _, v := range values {
		if max < v {
			max = v
		}
	}
	return max
}

func (items *ZArrayNumber[T]) Sum() T {
	sum := T(0)
	for _, v := range items.Get() {
		sum += v
	}
	return sum
}

/*
Avg calculates the average value of the elements in the ZArrayNumber.

It returns the average value as type T.
*/
func (items *ZArrayNumber[T]) Avg() T {
	return items.Sum() / T(items.Length())
}

func (items *ZArrayNumber[T]) SortAsc() *ZArrayNumber[T] {
	values := items.Sort(func(item1, item2 T, _, _ int) bool {
		return item1 < item2
	})
	return ZuiArrayNumber(values.Get())
}

func (items *ZArrayNumber[T]) SortDesc() *ZArrayNumber[T] {
	values := items.Sort(func(item1, item2 T, _, _ int) bool {
		return item1 > item2
	})
	return ZuiArrayNumber(values.Get())
}
