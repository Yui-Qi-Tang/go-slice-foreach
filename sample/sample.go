package main

import (
	"fmt"
	"github.com/Yui-Qi-Tang/go-slice-foreach/IntSlice"
)

func main() {
	testSlice := IntSlice{1,3,2,4,5,6}
	testSlice.forEach(func(x int) {
		fmt.Println(x)
	})
}