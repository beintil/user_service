package gosql

import "strings"

// { column_name data_type [ STORAGE { PLAIN | EXTERNAL | EXTENDED | MAIN | DEFAULT } ] [ COMPRESSION compression_method ] [ COLLATE collation ] [ column_constraint [ ... ] ] }
type column struct {
	// column name
	name string
	// data type
	dataType string
	// sort rule
	collate string
	// compression method
	compression string
	// column constraint
	constraint constraintColumn
	// column storage
	storage Storage
}

// Collate set sort rule
func (c *column) Collate(collation string) *column {
	c.collate = collation
	return c
}

// Compression set compression method
func (c *column) Compression(method string) *column {
	c.compression = method
	return c
}

// Name set column name
func (c *column) Name(name string) *column {
	c.name = name
	return c
}

// Type set column type
func (c *column) Type(dataType string) *column {
	c.dataType = dataType
	return c
}

// Storage set column storage
func (c *column) Storage(storage Storage) *column {
	c.storage = storage
	return c
}

// Constraint get constraint
func (c *column) Constraint() *constraintColumn {
	return &c.constraint
}

// IsEmpty check if empty
func (c *column) IsEmpty() bool {
	return c == nil || (c.name == "" && c.dataType == "" && c.collate == "" && c.compression == "" && c.constraint.IsEmpty())
}

// String render column
func (c *column) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if c.name != "" {
		b.WriteString(c.name)
	}
	if c.dataType != "" {
		b.WriteString(" " + c.dataType)
	}
	if c.compression != "" {
		b.WriteString(" COMPRESSION " + c.compression)
	}
	if c.collate != "" {
		b.WriteString(" COLLATE " + c.compression)
	}
	if c.storage != "" {
		b.WriteString(" STORAGE " + string(c.storage))
	}
	if !c.constraint.IsEmpty() {
		b.WriteString(c.constraint.String())
	}
	return b.String()
}
