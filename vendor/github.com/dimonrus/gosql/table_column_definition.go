package gosql

import "strings"

// list of column definition
type columnDefinitions []*columnDefinition

// String render all column definitions
func (c columnDefinitions) String() string {
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
func (c columnDefinitions) Len() int {
	return len(c)
}

// Add new definition
func (c *columnDefinitions) Add() (def *columnDefinition, n int) {
	// maloc new definition
	def = &columnDefinition{}
	// index of new definition in definition list
	n = len(*c)
	// append new definition in list
	*c = append(*c, def)
	return
}

// Swap definitions
func (c *columnDefinitions) Swap() {
	if c == nil || len(*c) == 0 {
		return
	}
	var i, j int
	j = len(*c) - 1
	for {
		if i == j || i-1 == j {
			break
		}
		buf := (*c)[i]
		(*c)[i] = (*c)[j]
		(*c)[j] = buf
		i++
		j--
	}
	return
}

// Remove definition by n
func (c *columnDefinitions) Remove(n int) *columnDefinitions {
	*c = append((*c)[:n], (*c)[n+1:]...)
	return c
}

// Clear remove all definitions
func (c *columnDefinitions) Clear() *columnDefinitions {
	*c = (*c)[:0]
	return c
}

// AddConstraint add constraint definition
func (c *columnDefinitions) AddConstraint() *constraintTable {
	def, _ := c.Add()
	return def.Constraint()
}

// AddColumn add column definition
func (c *columnDefinitions) AddColumn() *column {
	def, _ := c.Add()
	return def.Column()
}

// AddLike add like expression
func (c *columnDefinitions) AddLike() *likeTable {
	def, _ := c.Add()
	return def.Like()
}

// table column definition
// { column_name data_type [ COMPRESSION compression_method ] [ COLLATE collation ] [ column_constraint [ ... ] ]
//    | table_constraint
//    | LIKE source_table [ like_option ... ] }
type columnDefinition struct {
	// column
	column column
	// or constraint
	constraintTable constraintTable
	// or like expression
	like likeTable
}

// Column return column
func (d *columnDefinition) Column() *column {
	return &d.column
}

// Constraint return constraint
func (d *columnDefinition) Constraint() *constraintTable {
	return &d.constraintTable
}

// Like return like table
func (d *columnDefinition) Like() *likeTable {
	return &d.like
}

// IsEmpty check is columnDefinition is empty
func (d *columnDefinition) IsEmpty() bool {
	return d == nil || (d.column.IsEmpty() && d.constraintTable.IsEmpty() && d.like.IsEmpty())
}

// String render columnDefinition
func (d *columnDefinition) String() string {
	if d.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if !d.column.IsEmpty() {
		b.WriteString(d.column.String())
	} else if !d.constraintTable.IsEmpty() {
		b.WriteString(d.constraintTable.String())
	} else if !d.like.IsEmpty() {
		b.WriteString(d.like.String())
	}
	return b.String()
}
