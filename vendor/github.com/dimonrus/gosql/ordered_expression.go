package gosql

import (
	"bytes"
	"strings"
)

// orderedExpression helps to create ordered expression
type orderedExpression struct {
	// expression delimiter
	delimiter string
	// order
	expressions [][]byte
	// parameters
	params [][]any
	// builder
	builder strings.Builder
	// current iterator
	iterator int
	// bytes buffer
	buffer bytes.Buffer
}

// IsEmpty check if empty
func (a *orderedExpression) IsEmpty() bool {
	return a == nil || len(a.expressions) == 0 && len(a.params) == 0 && a.builder.Len() == 0
}

// Concat strings
func (a *orderedExpression) Concat(items ...string) []byte {
	var current = a.buffer.Len()
	for i := range items {
		a.buffer.WriteString(items[i])
	}
	return a.buffer.Bytes()[current:]
}

// Grow data. Multiplier is 32
func (a *orderedExpression) Grow(n int) *orderedExpression {
	a.builder.Grow(n * 32)
	buf := make([][]byte, 2*len(a.expressions)+n)
	copy(buf[0:], a.expressions)
	a.expressions = buf
	args := make([][]any, 2*len(a.params)+n)
	copy(args[0:], a.params)
	a.params = args
	return a
}

// Reset reset data
func (a *orderedExpression) Reset() *orderedExpression {
	a.builder.Reset()
	a.expressions = a.expressions[:0]
	a.params = a.params[:0]
	a.iterator = 0
	return a
}

// Iterator get iterator value
func (a *orderedExpression) Iterator() int {
	return a.iterator
}

// Add item data
func (a *orderedExpression) Add(order int, expression []byte, param ...any) *orderedExpression {
	if len(a.expressions) <= order {
		if order == 0 {
			a.Grow(4)
		} else {
			a.Grow(order * 2)
		}
	}
	a.expressions[order] = expression
	if len(param) > 0 {
		a.params[order] = param
	}
	return a
}

// Append item data
func (a *orderedExpression) Append(expression []byte, param ...any) *orderedExpression {
	a.Add(a.iterator, expression, param...)
	a.iterator++
	return a
}

// AppendArguments add argument
func (a *orderedExpression) AppendArguments(param ...any) *orderedExpression {
	if len(a.params) <= a.iterator {
		if a.iterator == 0 {
			a.Grow(4)
		} else {
			a.Grow(a.iterator * 2)
		}
	}
	if len(param) > 0 {
		a.params[a.iterator] = param
		a.iterator++
	}
	return a
}

// Delimiter set delimiter
func (a *orderedExpression) Delimiter(delimiter string) *orderedExpression {
	a.delimiter = delimiter
	return a
}

// GetArguments return params
func (a *orderedExpression) GetArguments() (params []any) {
	for i := range a.params {
		params = append(params, a.params[i]...)
	}
	return
}

// String render ordered expression
func (a *orderedExpression) String() string {
	if a.IsEmpty() {
		return ""
	}
	a.builder.Reset()
	if len(a.expressions) > 0 {
		// set default delimiter
		if a.delimiter == "" {
			a.delimiter = " "
		}
		var saved bool
		for i, s := range a.expressions {
			if len(s) == 0 {
				continue
			}
			if i > 0 && saved {
				a.builder.WriteString(a.delimiter)
			}
			a.builder.Write(s)
			saved = true
		}
	}
	return a.builder.String()
}
