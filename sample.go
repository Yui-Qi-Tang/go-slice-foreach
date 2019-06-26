package main

import (
	"fmt"

	"github.com/Yui-Qi-Tang/go-slice-foreach/slice"
)

var square slice.ANYMapF = func(v interface{}) interface{} {
	switch v.(type) {
	case int:
		return v.(int) * v.(int)
	default:
		return nil
	}
}

var notInt slice.ANYFilterF = func(v interface{}) bool {
	switch v.(type) {
	case int:
		return true
	default:
		return false
	}
}

func print(x interface{}) {
	fmt.Println(x)
}

func main() {

	testSlice := slice.IntSlice{1, 3, 2, 4, 5, 6}
	fmt.Println("Int slice forEach:", testSlice)
	testSlice.ForEach(func(x int) {
		fmt.Println("Just print =>", x)
	})

	testAny := slice.ANY{1, 3, "abc"}
	fmt.Println("ANY(interface{}):", testAny, "foreach")
	testAny.ForEach(print)

	integers := slice.ANY{1, 2, 3, 4, "gg", 5}
	fmt.Println("ANY(interface{}) slice Map(f: square(x) where x is int)", integers)
	// filter not int -> square for each element -> foreach print
	integers.Filter(notInt).Map(square).ForEach(print)

	// filter
	fmt.Println("Check if int in testAny", testAny, testAny.Filter(notInt))

}
