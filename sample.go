package main

import (
	"fmt"

	"github.com/Yui-Qi-Tang/go-slice-foreach/slice"
)

func main() {

	testSlice := slice.IntSlice{1, 3, 2, 4, 5, 6}
	fmt.Println("Int slice forEach:", testSlice)
	testSlice.ForEach(func(x int) {
		fmt.Println("Just print =>", x)
	})

	do := func(x interface{}) {
		fmt.Println(x)
	}

	testAny := slice.ANY{1, 3, "abc"}
	fmt.Println("ANY(interface{}):", testAny, "foreach")
	testAny.ForEach(do)

	fmt.Println("ANY(interface{}) slice Map(f: square(x) where x is int)")
	var square slice.ANYMapF = func(v interface{}) interface{} {
		switch v.(type) {
		case int:
			return v.(int) * v.(int)
		default:
			return nil
		}
	}

	integers := slice.ANY{1, 2, 3, 4, 5}
	// map
	newIntegers := integers.Map(square)
	// foreach
	newIntegers.ForEach(do)

	var checkIfInt slice.ANYFilterF = func(v interface{}) bool {
		switch v.(type) {
		case int:
			return true
		default:
			return false
		}
	}

	// filter
	fmt.Println("Check if int in testAny", testAny, testAny.Filter(checkIfInt))

}
