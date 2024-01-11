package main_test

import (
	"fmt"

	mzJson "github.com/chauhoangminhnguyen/material-zui-go/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Json", func() {
	Context("Decode", func() {
		It("default", func() {
			value := `{"key1": "value1", "key2": "value2"}`
			expected := map[string]interface{}{"key1": "value1", "key2": "value2"}

			result, err := mzJson.Decode(value, true)
			if err != nil {
				fmt.Errorf("Error decoding JSON: %v", err)
			}
			Expect(result).To(Equal(expected))
		})
	})
})
