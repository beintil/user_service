package gosql

import "strings"

// Delete query
type Delete struct {
	// with query
	with with
	// from
	from string
	// using
	using []string
	// where condition
	where Condition
	// returning
	returning expression
}

// IsEmpty check if query is empty
func (d *Delete) IsEmpty() bool {
	return d == nil || (d.with.Len() == 0 &&
		d.from == "" &&
		len(d.using) == 0 &&
		d.where.IsEmpty() &&
		d.returning.Len() == 0)
}

// From Set from value
func (d *Delete) From(from string) *Delete {
	d.from = from
	return d
}

// ResetFrom Set from empty string
func (d *Delete) ResetFrom() *Delete {
	d.from = ""
	return d
}

// Using add using
func (d *Delete) Using(using ...string) *Delete {
	d.using = append(d.using, using...)
	return d
}

// ResetUsing clear using
func (d *Delete) ResetUsing() *Delete {
	d.using = d.using[:0]
	return d
}

// Where set condition
func (d *Delete) Where() *Condition {
	return &d.where
}

// Returning Append returning expression
func (d *Delete) Returning() *expression {
	return &d.returning
}

// GetGetArguments get all values
func (d *Delete) GetGetArguments() []any {
	return append(append(d.with.GetArguments(), d.where.GetArguments()...))
}

// SQL Get sql query
func (d *Delete) SQL() (query string, params []any, returning []any) {
	return d.String(), d.GetGetArguments(), d.returning.GetArguments()
}

// String return result query
func (d *Delete) String() string {
	if d.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if d.with.Len() > 0 {
		b.WriteString(d.with.String() + " ")
	}
	b.WriteString("DELETE")
	if d.from != "" {
		b.WriteString(" FROM " + d.from)
	}
	if len(d.using) > 0 {
		b.WriteString(" USING " + strings.Join(d.using, ", "))
	}
	if !d.where.IsEmpty() {
		b.WriteString(" WHERE " + d.where.String())
	}
	if d.returning.Len() > 0 {
		b.WriteString(" RETURNING " + d.returning.String(", "))
	}
	b.WriteString(";")
	return b.String()
}

func NewDelete() *Delete {
	return &Delete{
		where: Condition{operator: ConditionOperatorAnd},
		with: with{
			keys: make(map[int]string),
		}}
}
