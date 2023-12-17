package gosql

import "strings"

//[ INCLUDE ( column_name [, ... ] ) ]
//[ WITH ( storage_parameter [ = value] [, ... ] ) ]
//[ USING INDEX TABLESPACE tablespace_name ]
type indexParameters struct {
	// include
	include expression
	// with
	with expression
	// using index tablespace
	tableSpace string
}

// With get with
func (i *indexParameters) With() *expression {
	return &i.with
}

// Include get include
func (i *indexParameters) Include() *expression {
	return &i.include
}

// TableSpace set tableSpace
func (i *indexParameters) TableSpace(tableSpace string) *indexParameters {
	i.tableSpace = tableSpace
	return i
}

// String render index parameters
func (i *indexParameters) String() string {
	if i.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if i.include.Len() > 0 {
		b.WriteString(" INCLUDE (" + i.include.String(", ") + ")")
	}
	if i.with.Len() > 0 {
		b.WriteString(" WITH (" + i.with.String(", ") + ")")
	}
	if i.tableSpace != "" {
		b.WriteString(" USING INDEX TABLESPACE " + i.tableSpace)
	}
	return b.String()
}

// IsEmpty is index parameter empty
func (i *indexParameters) IsEmpty() bool {
	return i == nil || (i.tableSpace == "" && i.include.Len() == 0 && i.with.Len() == 0)
}
