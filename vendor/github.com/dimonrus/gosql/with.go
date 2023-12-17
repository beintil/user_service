package gosql

import "strings"

// With SQL
type with struct {
	// map of with names
	keys map[int]string
	// slice of queries
	queries []*Select
	// recursive
	recursive bool
}

// Recursive query
func (w *with) Recursive() *with {
	w.recursive = true
	return w
}

// Len of with queries
func (w *with) Len() int {
	if w == nil {
		return 0
	}
	return len(w.keys)
}

// Get With
func (w *with) Get(name string) *Select {
	for j, key := range w.keys {
		if key == name {
			return w.queries[j]
		}
	}
	return nil
}

// Add With
func (w *with) Add(name string, qb *Select) *with {
	if name != "" && qb != nil {
		w.queries = append(w.queries, qb)
		if w.keys == nil {
			w.keys = make(map[int]string, 2)
		}
		w.keys[len(w.queries)-1] = name
	}
	return w
}

// Reset With query
func (w *with) Reset() *with {
	w.queries = w.queries[:0]
	w.keys = make(map[int]string, 2)
	w.recursive = false
	return w
}

// String for with
func (w *with) String() string {
	var b strings.Builder
	if w.Len() > 0 {
		b.WriteString("WITH ")
		if w.recursive {
			b.WriteString("RECURSIVE ")
		}
		for index, q := range w.queries {
			if index < len(w.queries)-1 {
				b.WriteString(w.keys[index] + " AS (" + q.String() + "),")
			} else {
				b.WriteString(w.keys[index] + " AS (" + q.String() + ")")
			}
		}
	}
	return b.String()
}

// GetArguments get values from all queries
func (w *with) GetArguments() []any {
	var params []any
	// order is important. map does not have an order
	if w.Len() > 0 {
		params = make([]any, 0, w.Len()*2)
		for index := range w.queries {
			params = append(params, w.queries[index].GetArguments()...)
		}
	}
	return params
}
