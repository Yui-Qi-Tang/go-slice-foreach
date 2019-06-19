package slice

// IntSlice a int slice
type IntSlice []int

func (i IntSlice) forEach(f func(v int)) {
	for _, v := range i {
		f(v)
	}
}
