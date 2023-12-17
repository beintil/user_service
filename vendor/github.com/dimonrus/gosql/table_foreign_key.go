package gosql

import "strings"

// FOREIGN KEY ( column_name [, ... ] ) REFERENCES reftable [ ( refcolumn [, ... ] ) ]
//
//	[ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ] [ ON DELETE referential_action ] [ ON UPDATE referential_action ] }
type foreignKey struct {
	// columns
	columns expression
	// reference
	referencesColumn referencesColumn
}

// Columns get columns
func (f *foreignKey) Columns() *expression {
	return &f.columns
}

// Columns get columns
func (f *foreignKey) Column(col ...string) *foreignKey {
	f.columns.Add(col...)
	return f
}

// References get references
func (f *foreignKey) References() *referencesColumn {
	return &f.referencesColumn
}

// String render foreign key
func (f *foreignKey) String() string {
	if f.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString(" FOREIGN KEY")
	if f.columns.Len() > 0 {
		b.WriteString(" (" + f.columns.String(", ") + ")")
	}
	if !f.referencesColumn.IsEmpty() {
		b.WriteString(" " + f.referencesColumn.String())
	}
	return b.String()
}

// IsEmpty is foreign key empty
func (f *foreignKey) IsEmpty() bool {
	return f == nil || (f.columns.Len() == 0 && f.referencesColumn.IsEmpty())
}
