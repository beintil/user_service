package gosql

import "strings"

// list of column definition
type ofDefinitions []*ofDefinition

// String render all column definitions
func (c ofDefinitions) String() string {
	b := strings.Builder{}
	for i, definition := range c {
		if i == 0 {
			b.WriteString(definition.String())
		} else {
			b.WriteString(", " + definition.String())
		}
	}
	return b.String()
}

// Len count of definition
func (c ofDefinitions) Len() int {
	return len(c)
}

// Add new definition
func (c *ofDefinitions) Add() (def *ofDefinition, n int) {
	// maloc new definition
	def = &ofDefinition{}
	// index of new definition in definition list
	n = len(*c)
	// append new definition in list
	*c = append(*c, def)
	return
}

// Remove definition by n
func (c *ofDefinitions) Remove(n int) *ofDefinitions {
	*c = append((*c)[:n], (*c)[n+1:]...)
	return c
}

// Clear remove all definitions
func (c *ofDefinitions) Clear() *ofDefinitions {
	*c = (*c)[:0]
	return c
}

// AddConstraint add constraint definition
func (c *ofDefinitions) AddConstraint() *constraintTable {
	def, _ := c.Add()
	return def.Constraint()
}

// AddColumn add column definition
func (c *ofDefinitions) AddColumn(name string) *ofColumn {
	def, _ := c.Add()
	return def.Column().Name(name)
}

//  OF type_name [ (
//  { column_name [ WITH OPTIONS ] [ column_constraint [ ... ] ]
//    | table_constraint }
//    [, ... ]
//) ]
type ofType struct {
	// of type
	name string
	// of definitions
	ofDefinitions ofDefinitions
}

// Name set name
func (t *ofType) Name(name string) *ofType {
	t.name = name
	return t
}

// Columns get columns definitions
func (t *ofType) Columns() *ofDefinitions {
	return &t.ofDefinitions
}

// IsEmpty check if empty
func (t *ofType) IsEmpty() bool {
	return t == nil || (t.name == "" && t.ofDefinitions.Len() == 0)
}

// String render column
func (t *ofType) String() string {
	if t.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if t.name != "" {
		b.WriteString(" OF " + t.name)
	}
	if t.ofDefinitions.Len() > 0 {
		b.WriteString(" (" + t.ofDefinitions.String() + ")")
	}
	return b.String()
}

//  { column_name [ WITH OPTIONS ] [ column_constraint [ ... ] ]
//    | table_constraint }
//    [, ... ]
type ofDefinition struct {
	// column definition
	ofColumn ofColumn
	// constrain definition
	constraintTable constraintTable
}

// Column get of column
func (c *ofDefinition) Column() *ofColumn {
	return &c.ofColumn
}

// Constraint get constraint
func (c *ofDefinition) Constraint() *constraintTable {
	return &c.constraintTable
}

// IsEmpty check if empty
func (c *ofDefinition) IsEmpty() bool {
	return c == nil || (c.ofColumn.IsEmpty() && c.constraintTable.IsEmpty())
}

// String render column
func (c *ofDefinition) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if !c.ofColumn.IsEmpty() {
		b.WriteString(c.ofColumn.String())
	} else if !c.constraintTable.IsEmpty() {
		b.WriteString(c.constraintTable.String())
	}
	return b.String()
}

// column_name [ WITH OPTIONS ] [ column_constraint [ ... ] ]
type ofColumn struct {
	// column name
	name string
	// WITH OPTIONS
	withOptions bool
	// constraint
	constraint constraintColumn
}

// Name set name
func (c *ofColumn) Name(name string) *ofColumn {
	c.name = name
	return c
}

// WithOptions reset name
func (c *ofColumn) WithOptions() *ofColumn {
	c.withOptions = true
	return c
}

// Constraint get constraint
func (c *ofColumn) Constraint() *constraintColumn {
	return &c.constraint
}

// IsEmpty check if empty
func (c *ofColumn) IsEmpty() bool {
	return c == nil || (c.name == "" && c.constraint.IsEmpty())
}

// String render column
func (c *ofColumn) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if c.name != "" {
		b.WriteString(c.name)
	}
	if c.withOptions {
		b.WriteString(" WITH OPTIONS")
	}
	if !c.constraint.IsEmpty() {
		b.WriteString(c.constraint.String())
	}
	return b.String()
}
