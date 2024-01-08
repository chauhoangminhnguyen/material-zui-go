package mz_utility

import (
	"fmt"
	"time"

	list "github.com/chauhoangminhnguyen/material-zui-go/list"
)

func UtilityTest() {
	now := time.Now()
	fmt.Println("UtilityTest", now)
	list.ListTest()
}
