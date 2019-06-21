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

//ANYMapF function template for any slice map
type ANYMapF func(v interface{}) interface{}

// ForEach comment for foreach
func (a ANY) ForEach(f func(v interface{})) {
	for _, v := range a {
		f(v)
	}
}

// Map do f on each element of ANY slice and return new ANY slice
func (a ANY) Map(f ANYMapF) (r ANY) {
	for _, v := range a {
		newV := f(v)
		r = append(r, newV)
	}

	return
}
