package slice

// IntSlice a int slice
type IntSlice []int

// ForEach executes a provided function once for each int slice element
func (i IntSlice) ForEach(f func(v int)) {
	for _, v := range i {
		f(v)
	}
}

// ANY slice ANY type
type ANY []interface{}

// ForEach comment for foreach
func (a ANY) ForEach(f func(v interface{})) {
	for _, v := range a {
		f(v)
	}
}
