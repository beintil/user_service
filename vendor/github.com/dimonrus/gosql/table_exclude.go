package gosql

import "strings"

// EXCLUDE [ USING index_method ] ( exclude_element WITH operator [, ... ] ) index_parameters [ WHERE ( predicate ) ] |
type excludeTable struct {
	// using index method
	using string
	// exclude element
	excludeElement excludeElement
	// with expression
	with expression
	// index parameters
	indexParameters indexParameters
	// where condition
	where *Condition
}

// Using set using
func (e *excludeTable) Using(using string) *excludeTable {
	e.using = using
	return e
}

// ExcludeElement return exclude element
func (e *excludeTable) ExcludeElement() *excludeElement {
	return &e.excludeElement
}

// With expression
func (e *excludeTable) With() *expression {
	return &e.with
}

// IndexParameters return index params
func (e *excludeTable) IndexParameters() *indexParameters {
	return &e.indexParameters
}

// Where condition
func (e *excludeTable) Where() *Condition {
	if e.where == nil {
		e.where = NewSqlCondition(ConditionOperatorAnd)
	}
	return e.where
}

// IsEmpty check if empty return true
func (e *excludeTable) IsEmpty() bool {
	return e == nil || (e.using == "" && e.excludeElement.IsEmpty() && e.with.Len() == 0 && e.indexParameters.IsEmpty() && e.where.IsEmpty())
}

// String render to string
func (e *excludeTable) String() string {
	if e.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString(" EXCLUDE")
	if e.using != "" {
		b.WriteString(" USING " + e.using)
	}
	if !e.excludeElement.IsEmpty() {
		if e.with.Len() > 0 {
			b.WriteString(" (" + e.excludeElement.String() + " WITH " + e.with.String(", ") + ")")
		} else {
			b.WriteString(" (" + e.excludeElement.String() + ")")
		}
	} else if e.with.Len() > 0 {
		b.WriteString(" WITH (" + e.with.String(", ") + ")")
	}
	if !e.indexParameters.IsEmpty() {
		b.WriteString(" " + e.indexParameters.String())
	}
	if !e.where.IsEmpty() {
		b.WriteString(" WHERE " + e.where.String())
	}
	return b.String()
}
