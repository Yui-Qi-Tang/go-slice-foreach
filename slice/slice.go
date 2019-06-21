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

//ANYMapF function signature for any slice map
type ANYMapF func(v interface{}) interface{}

//ANYFilterF function signature for any slice filter
type ANYFilterF func(v interface{}) bool

// ForEach comment for foreach
func (a ANY) ForEach(f func(v interface{})) {
	for _, v := range a {
		f(v)
	}
}

// Map do f on each element of ANY slice and return new ANY slice
func (a ANY) Map(f ANYMapF) (r ANY) {
	for _, v := range a {
		r = append(r, f(v))
	}

	return
}

// Filter return new ANY without f(v) is false
// need to return ANY?
func (a ANY) Filter(f ANYFilterF) (r ANY) {
	for _, v := range a {
		if f(v) {
			r = append(r, v)
		}
	}

	return
}

type ANYReduceF func(x, y interface{}) interface{}
type ANYRecusive func(x interface{}, y ANY) // Fold here?

func (a ANY) Reduce(f ANYReduceF) {
	// build a tree for folding?
}
