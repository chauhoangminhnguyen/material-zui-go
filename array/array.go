package mz_array

import (
	"fmt"
	"sort"
)

// type V any

type ZArray[T any] []T

func ZuiArray[T any](items []T) *ZArray[T] {
	result := &ZArray[T]{}
	for _, value := range items {
		*result = append(*result, value)
	}
	return result
}

func MapToZuiArray[T, V any](items ZArray[T], transform func(T) V) *ZArray[V] {
	return ZuiArray(Map(items, transform))
}

func (items *ZArray[T]) Get() []T {
	result := []T{}
	for _, value := range *items {
		result = append(result, value)
	}
	return result
}

func (items *ZArray[T]) At(index int) (T, error) {
	length := len(*items)
	if index < 0 {
		index = length + index
	}
	if index >= 0 && index < length {
		return (*items)[index], nil
	}
	return *new(T), fmt.Errorf("Index out of range")
}

func (items *ZArray[T]) AtWithDefault(index int, defaultValue T) T {
	value, err := items.At(index)
	if err != nil {
		return defaultValue
	}
	return value
}

func (items *ZArray[T]) AtByDefault(index int) T {
	value, _ := items.At(index)
	return value
}

func (items *ZArray[T]) Length() int {
	return len(*items)
}

func (items *ZArray[T]) LastIndex() int {
	return items.Length() - 1
}

func (items *ZArray[T]) First() T {
	return (*items).AtByDefault(0)
}

func (items *ZArray[T]) Last() T {
	return (*items).AtByDefault(items.LastIndex())
}

func (items *ZArray[T]) Append(i ...T) *ZArray[T] {
	*items = append(*items, i...)
	return items
}

func (items *ZArray[T]) Filter(transform func(T) bool) *ZArray[T] {
	results := ZArray[T]{}
	for _, v := range *items {
		if transform(v) {
			results.Append(v)
		}
	}
	return &results
}

func (items *ZArray[T]) Map(transform func(T) T) *ZArray[T] {
	rs := make(ZArray[T], len(*items))
	for i, v := range *items {
		rs[i] = transform(v)
	}
	return &rs
}

func (items *ZArray[T]) Some(condition func(T) bool) bool {
	for _, v := range *items {
		if condition(v) {
			return true
		}
	}
	return false
}

func (items *ZArray[T]) Every(condition func(T) bool) bool {
	for _, v := range *items {
		if condition(v) == false {
			return false
		}
	}
	return true
}

func (items *ZArray[T]) Sort(sortFnc func(item1, item2 T, idx1, idx2 int) bool) *ZArray[T] {
	values := items.Get()
	sort.SliceStable(values, func(i, j int) bool {
		return sortFnc(values[i], values[j], i, j)
	})
	return ZuiArray(values)
}
