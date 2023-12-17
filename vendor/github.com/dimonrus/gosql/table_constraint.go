package gosql

import "strings"

// [ CONSTRAINT constraint_name ]
// { CHECK ( expression ) [ NO INHERIT ] |
//
//	UNIQUE [ NULLS [ NOT ] DISTINCT ] ( column_name [, ... ] ) index_parameters |
//	PRIMARY KEY ( column_name [, ... ] ) index_parameters |
//	EXCLUDE [ USING index_method ] ( exclude_element WITH operator [, ... ] ) index_parameters [ WHERE ( predicate ) ] |
//	FOREIGN KEY ( column_name [, ... ] ) REFERENCES reftable [ ( refcolumn [, ... ] ) ]
//	  [ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ] [ ON DELETE referential_action ] [ ON UPDATE referential_action ] }
//
// [ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
type constraintTable struct {
	// name
	name string
	// check expression
	check *Condition
	// unique index
	unique columnIndexParameters
	// primary key
	primary columnIndexParameters
	// exclude
	exclude excludeTable
	// foreign key
	foreignKey foreignKey
	// no inherit
	noInherit bool
	// deferrable
	deferrable *bool
	// initially
	initially string
	// nulls not distinct
	nullsNotDistinct bool
}

// Name set name
func (c *constraintTable) Name(name string) *constraintTable {
	c.name = name
	return c
}

// Check get check expression
func (c *constraintTable) Check() *Condition {
	if c.check == nil {
		c.check = NewSqlCondition(ConditionOperatorAnd)
	}
	return c.check
}

// NoInherit set no inherit
func (c *constraintTable) NoInherit() *constraintTable {
	c.noInherit = true
	return c
}

// NullNotDistinct is unique constraint null not distinct
func (c *constraintTable) NullNotDistinct() *constraintTable {
	c.nullsNotDistinct = true
	return c
}

// Unique get unique expression
func (c *constraintTable) Unique() *columnIndexParameters {
	return &c.unique
}

// PrimaryKey get primary key expression
func (c *constraintTable) PrimaryKey() *columnIndexParameters {
	return &c.primary
}

// Exclude get exclude expression
func (c *constraintTable) Exclude() *excludeTable {
	return &c.exclude
}

// ForeignKey get foreign key expression
func (c *constraintTable) ForeignKey() *foreignKey {
	return &c.foreignKey
}

// Deferrable set deferrable
func (c *constraintTable) Deferrable(deferrable *bool) *constraintTable {
	c.deferrable = deferrable
	return c
}

// Initially set initially
func (c *constraintTable) Initially(initially string) *constraintTable {
	c.initially = initially
	return c
}

// IsEmpty check if constraint is empty
func (c *constraintTable) IsEmpty() bool {
	return c == nil || (c.name == "" &&
		c.check.IsEmpty() &&
		c.unique.IsEmpty() &&
		c.primary.IsEmpty() &&
		c.exclude.IsEmpty() &&
		c.foreignKey.IsEmpty() &&
		c.deferrable == nil &&
		c.initially == "" &&
		!c.nullsNotDistinct)
}

// String render table constraint
func (c *constraintTable) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if c.name != "" {
		b.WriteString(" CONSTRAINT " + c.name)
	}
	if !c.check.IsEmpty() {
		b.WriteString(" CHECK " + c.check.String())
	}
	if c.noInherit {
		b.WriteString(" NO INHERIT")
	}
	if !c.unique.IsEmpty() {
		if c.nullsNotDistinct {
			b.WriteString(" UNIQUE NULLS NOT DISTINCT " + c.unique.String())
		} else {
			b.WriteString(" UNIQUE " + c.unique.String())
		}
	}
	if !c.primary.IsEmpty() {
		b.WriteString(" PRIMARY KEY " + c.primary.String())
	}
	if !c.exclude.IsEmpty() {
		b.WriteString(c.exclude.String())
	}
	if !c.foreignKey.IsEmpty() {
		b.WriteString(c.foreignKey.String())
	}
	if c.deferrable != nil {
		if *c.deferrable {
			b.WriteString(" DEFERRABLE")
		} else {
			b.WriteString(" NOT DEFERRABLE")
		}
	}
	if c.initially != "" {
		b.WriteString(" INITIALLY " + c.initially)
	}
	return b.String()[1:]
}
