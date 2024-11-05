package mz_map

import (
	"fmt"
	"sort"
)

// Keys returns the keys of a given map[string]V.
//
// It takes a map[string]V as a parameter and returns a slice of strings that contains all the keys
// of the map. The keys are sorted in ascending order.
//
// Params:
// - value: a map[string]V to retrieve the keys from.
//
// Returns:
// - []string: a slice of strings containing all the keys of the input map, sorted in ascending order.
func Keys[V any](value map[string]V) []string {
	keys := make([]string, 0, len(value))
	for key := range value { // the loop of map object is random the order
		fmt.Println(key)
		keys = append(keys, key)
	}
	sort.Strings(keys) // sort keys
	return keys
}

func Values[V any](value map[string]V) []V {
	values := make([]V, 0, len(value))
	for _, v := range value {
		values = append(values, v)
	}
	return values
}

func Entries[V interface{}](value map[string]V) [][2]interface{} {
	entries := make([][2]interface{}, 0, len(value))
	for k, v := range value {
		entries = append(entries, [2]interface{}{k, v})
	}
	return entries
}

/*
SortKey sorts the keys of a map[string]V and returns a new map with the sorted keys.

value: the map[string]V to be sorted

Returns: a new map with the sorted keys
*/
func SortKey[V any](value map[string]V) map[string]V {
	result := map[string]V{}
	keys := Keys(value)
	// sort.Strings(keys) // sort keys
	for _, k := range keys {
		result[k] = value[k]
	}
	return result
}
