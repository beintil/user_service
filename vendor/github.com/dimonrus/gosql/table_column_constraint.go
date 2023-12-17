package gosql

import "strings"

// [ CONSTRAINT constraint_name ]
// { NOT NULL |
//
//	NULL |
//	CHECK ( expression ) [ NO INHERIT ] |
//	DEFAULT default_expr |
//	GENERATED ALWAYS AS ( generation_expr ) STORED |
//	GENERATED { ALWAYS | BY DEFAULT } AS IDENTITY [ ( sequence_options ) ] |
//	UNIQUE [ NULLS [ NOT ] DISTINCT ] index_parameters |
//	PRIMARY KEY index_parameters |
//	REFERENCES reftable [ ( refcolumn ) ] [ MATCH FULL | MATCH PARTIAL | MATCH SIMPLE ]
//	  [ ON DELETE referential_action ] [ ON UPDATE referential_action ] }
//
// [ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
type constraintColumn struct {
	// name
	name string
	// nullable
	notNull bool
	// check expression
	check detailedExpression
	// default
	def string
	// generated always as expression
	generatedAlwaysAs expression
	// generated
	generated detailedExpression
	// unique index
	unique *indexParameters
	// primary key
	primary *indexParameters
	// references
	references referencesColumn
	// deferrable
	deferrable *bool
	// initially
	initially string
	// nulls not distinct
	nullsNotDistinct bool
}

// Name set name
func (c *constraintColumn) Name(name string) *constraintColumn {
	c.name = name
	return c
}

// NotNull is constraint nullable
func (c *constraintColumn) NotNull() *constraintColumn {
	c.notNull = true
	return c
}

// NullNotDistinct is unique constraint null not distinct
func (c *constraintColumn) NullNotDistinct() *constraintColumn {
	c.nullsNotDistinct = true
	return c
}

// Check detailed expression
func (c *constraintColumn) Check() *detailedExpression {
	return &c.check
}

// Default set default
func (c *constraintColumn) Default(def string) *constraintColumn {
	c.def = def
	return c
}

// GeneratedAlways expression
func (c *constraintColumn) GeneratedAlways() *expression {
	return &c.generatedAlwaysAs
}

// Generated expression
func (c *constraintColumn) Generated() *detailedExpression {
	return &c.generated
}

// Unique set unique
func (c *constraintColumn) Unique() *indexParameters {
	c.unique = &indexParameters{}
	return c.unique
}

// PrimaryKey set primary
func (c *constraintColumn) PrimaryKey() *indexParameters {
	c.primary = &indexParameters{}
	return c.primary
}

// References get references
func (c *constraintColumn) References() *referencesColumn {
	return &c.references
}

// Deferrable set deferrable
func (c *constraintColumn) Deferrable(deferrable *bool) *constraintColumn {
	c.deferrable = deferrable
	return c
}

// Initially set initially
func (c *constraintColumn) Initially(initially string) *constraintColumn {
	c.initially = initially
	return c
}

// String render column constraint
func (c *constraintColumn) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if c.name != "" {
		b.WriteString(" CONSTRAINT " + c.name)
	}
	if c.notNull {
		b.WriteString(" NOT NULL")
	}
	if !c.check.IsEmpty() {
		b.WriteString(" CHECK " + c.check.String())
	}
	if c.def != "" {
		b.WriteString(" DEFAULT " + c.def)
	}
	if c.generatedAlwaysAs.Len() > 0 {
		b.WriteString(" GENERATED ALWAYS AS (" + c.generatedAlwaysAs.String(", ") + ") STORED")
	}
	if !c.generated.IsEmpty() {
		b.WriteString(" GENERATED " + c.generated.GetDetail() + " AS IDENTITY")
		if c.generated.Expression().Len() > 0 {
			b.WriteString(" " + c.generated.Expression().String(", "))
		}
	}
	if c.unique != nil {
		if c.nullsNotDistinct {
			b.WriteString(" UNIQUE NULLS NOT DISTINCT" + c.unique.String())
		} else {
			b.WriteString(" UNIQUE" + c.unique.String())
		}
	}
	if c.primary != nil {
		b.WriteString(" PRIMARY KEY" + c.primary.String())
	}
	if !c.references.IsEmpty() {
		b.WriteString(" " + c.references.String())
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
	return b.String()
}

// IsEmpty check if empty
func (c *constraintColumn) IsEmpty() bool {
	return c == nil || (c.name == "" &&
		!c.notNull &&
		c.check.IsEmpty() &&
		c.def == "" &&
		c.generatedAlwaysAs.Len() == 0 &&
		c.generated.IsEmpty() &&
		c.unique == nil &&
		c.primary == nil &&
		c.references.IsEmpty() &&
		c.deferrable == nil &&
		c.initially == "" &&
		!c.nullsNotDistinct)
}
