package gosql

import "strings"

// references column
type referencesColumn struct {
	// target table
	target string
	// target columns
	column expression
	// match full partial simple
	match string
	// on update
	update string
	// on delete
	delete string
}

// RefTable set reference table
func (r *referencesColumn) RefTable(table string) *referencesColumn {
	r.target = table
	return r
}

// Columns reference columns
func (r *referencesColumn) Columns() *expression {
	return &r.column
}

// Column reference columns
func (r *referencesColumn) Column(col ...string) *referencesColumn {
	r.column.Add(col...)
	return r
}

// Match set match
func (r *referencesColumn) Match(match string) *referencesColumn {
	r.match = match
	return r
}

// OnUpdate set on update
func (r *referencesColumn) OnUpdate(update string) *referencesColumn {
	r.update = update
	return r
}

// OnDelete set on delete
func (r *referencesColumn) OnDelete(delete string) *referencesColumn {
	r.delete = delete
	return r
}

// IsEmpty check if empty
func (r *referencesColumn) IsEmpty() bool {
	return r == nil || (r.column.Len() == 0 &&
		r.target == "" &&
		r.update == "" &&
		r.delete == "" &&
		r.match == "")
}

// String create reference for column
func (r *referencesColumn) String() string {
	if r.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("REFERENCES " + r.target)
	if r.column.Len() > 0 {
		b.WriteString(" (" + r.column.String(", ") + ")")
	}
	if r.match != "" {
		b.WriteString(" MATCH " + r.match)
	}
	if r.delete != "" {
		b.WriteString(" ON DELETE " + r.delete)
	}
	if r.update != "" {
		b.WriteString(" ON UPDATE " + r.update)
	}
	return b.String()
}
