package main

import (
	"fmt"

	"github.com/Yui-Qi-Tang/go-slice-foreach/slice"
)

func main() {
	testSlice := slice.IntSlice{1, 3, 2, 4, 5, 6}
	testSlice.ForEach(func(x int) {
		fmt.Println(x)
	})
}
