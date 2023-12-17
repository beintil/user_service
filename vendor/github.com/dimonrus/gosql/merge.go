package gosql

import "strings"

// [ WITH with_query [, ...] ]
// MERGE INTO [ ONLY ] target_table_name [ * ] [ [ AS ] target_alias ]
// USING data_source ON join_condition
// when_clause [...]
//
// where data_source is:
//
// { [ ONLY ] source_table_name [ * ] | ( source_query ) } [ [ AS ] source_alias ]
//
// and when_clause is:
//
// { WHEN MATCHED [ AND condition ] THEN { merge_update | merge_delete | DO NOTHING } |
//
//	WHEN NOT MATCHED [ AND condition ] THEN { merge_insert | DO NOTHING } }
//
// and merge_insert is:
//
// INSERT [( column_name [, ...] )]
// [ OVERRIDING { SYSTEM | USER } VALUE ]
// { VALUES ( { expression | DEFAULT } [, ...] ) | DEFAULT VALUES }
//
// and merge_update is:
//
// UPDATE SET { column_name = { expression | DEFAULT } |
//
//	( column_name [, ...] ) = ( { expression | DEFAULT } [, ...] ) } [, ...]
//
// and merge_delete is:
//
// DELETE
type Merge struct {
	// WITH queries
	with with
	// INTO
	into detailedExpression
	// Using
	using string
	// When clause
	when []when
}

// With queries
func (m *Merge) With() *with {
	return &m.with
}

// Into target
func (m *Merge) Into(target string) *Merge {
	m.into.SetDetail(target)
	return m
}

// IntoOnly target
func (m *Merge) IntoOnly(target string) *Merge {
	m.into.SetDetail("ONLY")
	return m.Into(target)
}

// Using datasource
func (m *Merge) Using(datasource string) *Merge {
	m.using = datasource
	return m
}

// UsingOnly datasource
func (m *Merge) UsingOnly(datasource string) *Merge {
	return m.Using("ONLY " + datasource)
}

// When clause
func (m *Merge) When() *when {
	m.when = append(m.when, when{})
	return &m.when[len(m.when)-1]
}

// isWhenEmpty is when is empty
func (m *Merge) isWhenEmpty() bool {
	for _, w := range m.when {
		if !w.IsEmpty() {
			return false
		}
	}
	return true
}

// isWhenEmpty is when is empty
func (m *Merge) whenString() string {
	b := strings.Builder{}
	for _, w := range m.when {
		b.WriteString(w.String())
	}
	return b.String()
}

// isWhenEmpty is when is empty
func (m *Merge) whenArguments() []any {
	args := make([]any, 0)
	for i := range m.when {
		args = append(args, m.when[i].GetArguments()...)
	}
	return args
}

// IsEmpty if empty
func (m *Merge) IsEmpty() bool {
	return m == nil ||
		m.with.Len() == 0 &&
			m.into.IsEmpty() &&
			len(m.using) == 0 &&
			m.isWhenEmpty()
}

// String build query
func (m *Merge) String() string {
	b := strings.Builder{}
	// With render
	if m.with.Len() > 0 {
		b.WriteString(m.with.String() + " ")
	}
	if !m.into.IsEmpty() {
		b.WriteString("MERGE INTO ")
		if len(m.into.GetDetail()) > 0 {
			b.WriteString(m.into.GetDetail() + " ")
		}
		b.WriteString(m.into.String())
	}
	if len(m.using) > 0 {
		b.WriteString("USING " + m.using)
	}
	if !m.isWhenEmpty() {
		b.WriteString(m.whenString())
	}
	return b.String() + ";"
}

// GetArguments Get merge arguments
func (m *Merge) GetArguments() []any {
	return append(m.with.GetArguments(), m.whenArguments()...)
}

// SQL Get sql query
func (m *Merge) SQL() (query string, params []any, returning []any) {
	return m.String(), m.GetArguments(), nil
}

// NewMerge merge constructor
func NewMerge() *Merge {
	return &Merge{}
}

type when struct {
	// condition
	condition Condition
	// update
	update expression
	// delete
	delete expression
	// insert
	insert mergeInsert
	// do nothing
	doNoting bool
}

// Condition set condition
func (w *when) Condition() *Condition {
	return &w.condition
}

// Update set merge update
func (w *when) Update() *expression {
	return &w.update
}

// Delete set merge delete
func (w *when) Delete() *when {
	w.delete.Add("DELETE")
	return w
}

// Insert set merge delete
func (w *when) Insert() *mergeInsert {
	return &w.insert
}

// DoNothing set do nothing
func (w *when) DoNothing() *when {
	w.doNoting = true
	return w
}

// String build query
func (w *when) String() string {
	if w.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if w.update.Len() > 0 {
		b.WriteString(" WHEN MATCHED ")
		if !w.condition.IsEmpty() {
			b.WriteString("AND " + w.Condition().String() + " ")
		}
		if w.doNoting {
			b.WriteString("THEN DO NOTHING")
		} else {
			b.WriteString("THEN UPDATE SET " + w.update.String(", "))
		}
	} else if w.delete.Len() > 0 {
		b.WriteString(" WHEN MATCHED ")
		if !w.condition.IsEmpty() {
			b.WriteString("AND " + w.Condition().String())
		}
		b.WriteString("THEN " + w.delete.String(","))
	} else if !w.insert.IsEmpty() {
		b.WriteString(" WHEN NOT MATCHED ")
		if !w.condition.IsEmpty() {
			b.WriteString("AND " + w.Condition().String() + " ")
		}
		if w.doNoting {
			b.WriteString("THEN DO NOTHING")
		} else {
			b.WriteString("THEN " + w.insert.String())
		}
	}
	return b.String()
}

// IsEmpty check if empty
func (w *when) IsEmpty() bool {
	return w == nil ||
		w.condition.IsEmpty() &&
			w.insert.IsEmpty() &&
			w.update.Len() == 0 &&
			w.delete.Len() == 0
}

// GetArguments get all arguments
func (w *when) GetArguments() []any {
	if w.update.ArgLen() > 0 {
		return append(w.condition.GetArguments(), w.update.GetArguments()...)
	} else {
		return append(w.condition.GetArguments(), w.insert.GetArguments()...)
	}
}

type mergeInsert struct {
	// Columns
	columns expression
	// OVERRIDING
	overriding string
	// VALUES
	values detailedExpression
}

// Columns add columns
func (m *mergeInsert) Columns(columns ...string) *mergeInsert {
	m.columns.Add(columns...)
	return m
}

// Overriding overriding
func (m *mergeInsert) Overriding(overriding string) *mergeInsert {
	m.overriding = overriding
	return m
}

// Values set values
func (m *mergeInsert) Values() *expression {
	return &m.values.expression
}

// DefaultValues set default values
func (m *mergeInsert) DefaultValues() *mergeInsert {
	m.values.SetDetail("DEFAULT VALUES")
	return m
}

// String build query
func (m *mergeInsert) String() string {
	if m.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("INSERT ")
	// columns
	if m.columns.Len() > 0 {
		b.WriteString("(" + m.columns.String(", ") + ") ")
	}
	if len(m.overriding) > 0 {
		b.WriteString(" OVERRIDING " + m.overriding)
	}
	if !m.values.IsEmpty() {
		if len(m.values.GetDetail()) > 0 {
			b.WriteString("VALUES " + m.values.GetDetail())
		} else {
			b.WriteString("VALUES (" + m.values.Expression().String(", ") + ")")
		}
	}
	return b.String()
}

// IsEmpty check if empty
func (m *mergeInsert) IsEmpty() bool {
	return m == nil ||
		m.columns.Len() == 0 &&
			m.values.IsEmpty()
}

// GetArguments get all arguments
func (m *mergeInsert) GetArguments() []any {
	return append(m.columns.GetArguments(), m.values.Expression().GetArguments()...)
}
