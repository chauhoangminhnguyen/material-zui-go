package mz_utility

import (
	"fmt"
	"time"

	array "github.com/chauhoangminhnguyen/material-zui-go/array"
)

func UtilityTest() {
	now := time.Now()
	fmt.Println("UtilityTest", now)
	array.ArrayTest()
}
