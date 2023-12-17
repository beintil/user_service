package gosql

import (
	"strings"
)

const (
	// ConflictActionNothing On conflict action do nothing
	ConflictActionNothing = "NOTHING"
	// ConflictActionUpdate On conflict action do update nothing
	ConflictActionUpdate = "UPDATE"
)

// Insert query builder
type Insert struct {
	// with query
	with with
	// into table name
	into string
	// from insert
	from []string
	// list of columns
	columns expression
	// conflict expression
	conflict conflict
	// returning
	returning expression
}

// Get sql insert query
func (i *Insert) String() string {
	b := strings.Builder{}
	if i.with.Len() > 0 {
		b.WriteString(i.with.String() + " ")
	}
	if i.into != "" {
		b.WriteString("INSERT INTO " + i.into)
	}
	if i.columns.Len() > 0 {
		b.WriteString(" (" + i.columns.String(", ") + ")")
	}
	if len(i.from) > 0 {
		b.WriteString(" " + strings.Join(i.from, ", "))
	} else if i.columns.ArgLen() > 0 {
		b.WriteString(" VALUES ")
		cols := len(i.columns.Split())
		for j := 1; j <= i.columns.ArgLen(); {
			b.WriteString("(")
			for u := 0; u < cols; u++ {
				b.WriteString("?")
				j++
				if u == cols-1 {
					break
				}
				b.WriteString(", ")
			}
			b.WriteString(")")
			if j >= i.columns.ArgLen() {
				break
			}
			b.WriteString(", ")
		}
	}
	if !i.conflict.IsEmpty() {
		b.WriteString(" " + i.Conflict().String())
	}
	if i.returning.Len() > 0 {
		b.WriteString(" RETURNING " + i.returning.String(", "))
	}
	b.WriteString(";")
	return b.String()
}

// IsEmpty return true if all parts not filed
func (i *Insert) IsEmpty() bool {
	return i == nil || (i.with.Len() == 0 &&
		i.into == "" &&
		len(i.from) == 0 &&
		i.columns.Len() == 0 &&
		i.columns.ArgLen() == 0 &&
		i.returning.Len() == 0 &&
		i.conflict.IsEmpty())
}

// SQL Get sql query
func (i *Insert) SQL() (query string, params []any, returning []any) {
	return i.String(), i.GetArguments(), i.returning.GetArguments()
}

// GetArguments get all arguments
func (i *Insert) GetArguments() []any {
	return append(append(i.with.GetArguments(), i.columns.GetArguments()...), i.conflict.GetArguments()...)
}

// SetConflict set conflict
func (i *Insert) SetConflict(conflict conflict) *Insert {
	i.conflict = conflict
	return i
}

// Conflict get conflict expression
func (i *Insert) Conflict() *conflict {
	return &i.conflict
}

// ResetConflict reset conflict expression
func (i *Insert) ResetConflict() *Insert {
	i.conflict = conflict{}
	return i
}

// Returning Get returning expression
func (i *Insert) Returning() *expression {
	return &i.returning
}

// From insert from
func (i *Insert) From(from ...string) *Insert {
	i.from = append(i.from, from...)
	return i
}

// ResetFrom clear from
func (i *Insert) ResetFrom() *Insert {
	i.from = i.from[:0]
	return i
}

// Columns get columns for insert
func (i *Insert) Columns() *expression {
	return &i.columns
}

// Into Set into value
func (i *Insert) Into(into string) *Insert {
	i.into = into
	return i
}

// ResetInto Set into empty string
func (i *Insert) ResetInto() *Insert {
	i.into = ""
	return i
}

// With Get with query
func (i *Insert) With() *with {
	return &i.with
}

// NewInsert new insert query builder
func NewInsert() *Insert {
	return &Insert{
		with: with{
			keys: make(map[int]string),
		},
	}
}
