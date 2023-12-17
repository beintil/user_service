package gosql

import "strings"

// EnumDelimiter for join strings
const EnumDelimiter = "#"

// Expression slice
type expression struct {
	// list of expressions
	list strings.Builder
	// params
	params []any
}

// Append expressions
func (e *expression) Append(expr string, args ...any) {
	e.Add(expr)
	e.Arg(args...)
}

// Len of expressions
func (e *expression) Len() int {
	if e == nil {
		return 0
	}
	return e.list.Len()
}

// ArgLen len of arguments
func (e *expression) ArgLen() int {
	if e == nil {
		return 0
	}
	return len(e.params)
}

// Reset expressions
func (e *expression) Reset() {
	if e.Len() == 0 {
		return
	}
	e.list.Reset()
	e.params = e.params[:0]
}

// String all expressions
func (e *expression) String(delimiter string) string {
	if e.Len() == 0 {
		return ""
	}
	return strings.ReplaceAll(e.list.String(), EnumDelimiter, delimiter)
}

// Arg add params
func (e *expression) Arg(args ...any) {
	e.params = append(e.params, args...)
	return
}

// Add expression items
func (e *expression) Add(item ...string) {
	for _, arg := range item {
		if e.Len() == 0 {
			e.list.WriteString(arg)
		} else {
			e.list.WriteString(EnumDelimiter + arg)
		}
	}
	return
}

// GetArguments return params
func (e *expression) GetArguments() []any {
	if e == nil {
		return nil
	}
	return e.params
}

// Get expressions
func (e *expression) Get(delimiter string) (string, []any) {
	return e.String(delimiter), e.GetArguments()
}

// Grow memory. Multiplier is 32
func (e *expression) Grow(n int) *expression {
	e.list.Grow(n * 32)
	args := make([]any, 2*len(e.params)+n)
	copy(args[0:], e.params)
	e.params = args[:len(e.params)]
	return e
}

// Split to slice of string
func (e *expression) Split() []string {
	return strings.SplitN(e.list.String(), EnumDelimiter, -1)
}

// NewExpression init expression
func NewExpression() *expression {
	return (&expression{}).Grow(4)
}
