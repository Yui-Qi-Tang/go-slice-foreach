package slice

// IntSlice a int slice
type IntSlice []int

// ForEach executes a provided function once for each int slice element
func (i IntSlice) ForEach(f func(v int)) {
	for _, v := range i {
		f(v)
	}
}
