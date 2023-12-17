package gosql

import "strings"

// Update query builder
type Update struct {
	// with query
	with with
	// table name
	table string
	// set of changes
	set expression
	// from source
	from []string
	// condition
	where Condition
	// returning
	returning expression
}

// IsEmpty check if query is empty
func (u *Update) IsEmpty() bool {
	return u == nil || (u.with.Len() == 0 &&
		u.table == "" &&
		len(u.from) == 0 &&
		u.where.IsEmpty() &&
		u.set.Len() == 0 &&
		u.returning.Len() == 0)
}

// From update
func (u *Update) From(from ...string) *Update {
	u.from = append(u.from, from...)
	return u
}

// ResetFrom clear from
func (u *Update) ResetFrom() *Update {
	u.from = u.from[:0]
	return u
}

// Where set condition
func (u *Update) Where() *Condition {
	return &u.where
}

// String return result query
func (u *Update) String() string {
	if u.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if u.with.Len() > 0 {
		b.WriteString(u.with.String() + " ")
	}
	if u.table != "" {
		b.WriteString("UPDATE " + u.table)
	}
	if u.set.Len() > 0 {
		b.WriteString(" SET " + u.set.String(", "))
	}
	if len(u.from) != 0 {
		b.WriteString(" FROM " + strings.Join(u.from, ", "))
	}
	if !u.where.IsEmpty() {
		b.WriteString(" WHERE " + u.where.String())
	}
	if u.returning.Len() > 0 {
		b.WriteString(" RETURNING " + u.returning.String(", "))
	}
	b.WriteString(";")
	return b.String()
}

// Set get set
func (u *Update) Set() *expression {
	return &u.set
}

// Table Set table
func (u *Update) Table(table string) *Update {
	u.table = table
	return u
}

// ResetTable reset table
func (u *Update) ResetTable() *Update {
	u.table = ""
	return u
}

// GetArguments get all values
func (u *Update) GetArguments() []any {
	return append(append(u.with.GetArguments(), u.set.GetArguments()...), u.where.GetArguments()...)
}

// Returning get returning expression
func (u *Update) Returning() *expression {
	return &u.returning
}

// With Append with query
func (u *Update) With() *with {
	return &u.with
}

// SQL Get sql query
func (u *Update) SQL() (query string, params []any, returning []any) {
	return u.String(), append(append(u.with.GetArguments(), u.set.GetArguments()...), u.where.GetArguments()...), u.returning.GetArguments()
}

// NewUpdate Update Query Builder
func NewUpdate() *Update {
	return &Update{
		where: Condition{operator: ConditionOperatorAnd},
		with: with{
			keys: make(map[int]string),
		},
	}
}
