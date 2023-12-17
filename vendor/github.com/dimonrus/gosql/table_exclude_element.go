package gosql

import "strings"

// exclude element
// { column_name | ( expression ) } [ opclass ] [ ASC | DESC ] [ NULLS { FIRST | LAST } ]
type excludeElement struct {
	// column name
	column string
	// expression
	expression expression
	// operator class
	opclass string
	// direction
	direction string
	// nulls expression
	nulls *string
}

// Column set column name
func (e *excludeElement) Column(column string) *excludeElement {
	e.column = column
	return e
}

// OpClass set opclass
func (e *excludeElement) OpClass(opclass string) *excludeElement {
	e.opclass = opclass
	return e
}

// Direction set direction
func (e *excludeElement) Direction(direction string) *excludeElement {
	e.direction = direction
	return e
}

// Nulls set nulls
func (e *excludeElement) Nulls(nulls *string) *excludeElement {
	e.nulls = nulls
	return e
}

// Expression get expression
func (e *excludeElement) Expression() *expression {
	return &e.expression
}

// String render exclude element
func (e *excludeElement) String() string {
	if e.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if e.column != "" {
		b.WriteString(e.column)
	} else if e.expression.Len() > 0 {
		b.WriteString("(" + e.expression.String(", ") + ")")
	}
	if e.opclass != "" {
		b.WriteString(" " + e.opclass)
	}
	if e.direction != "" {
		b.WriteString(" " + e.direction)
	}
	if e.nulls != nil {
		b.WriteString(" NULLS " + *e.nulls)
	}
	return b.String()
}

// IsEmpty is exclude element is empty
func (e *excludeElement) IsEmpty() bool {
	return e == nil || (e.column == "" && e.expression.Len() == 0 && e.opclass == "" && e.direction == "" && e.nulls == nil)
}
