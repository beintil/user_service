package gosql

// where add action is one of:
//
// DROP [ COLUMN ] [ IF EXISTS ] column_name [ RESTRICT | CASCADE ]
// DROP DEFAULT
// DROP NOT NULL
// DROP EXPRESSION [ IF EXISTS ]
// DROP IDENTITY [ IF EXISTS ]
// DROP CONSTRAINT [ IF EXISTS ]  constraint_name [ RESTRICT | CASCADE ]
type alterTableActionDrop struct {
	// ordered expression
	ordered orderedExpression
}

// Column specify column name
func (a *alterTableActionDrop) Column(name string) *alterTableActionDrop {
	a.ordered.Add(0, a.ordered.Concat("DROP COLUMN"))
	a.ordered.Add(2, a.ordered.Concat(name))
	return a
}

// IfExists add if exists
func (a *alterTableActionDrop) IfExists() *alterTableActionDrop {
	a.ordered.Add(1, a.ordered.Concat("IF EXISTS"))
	return a
}

// Restrict drop
func (a *alterTableActionDrop) Restrict() *alterTableActionDrop {
	a.ordered.Add(3, a.ordered.Concat("RESTRICT"))
	return a
}

// Cascade drop
func (a *alterTableActionDrop) Cascade() *alterTableActionDrop {
	a.ordered.Add(3, a.ordered.Concat("CASCADE"))
	return a
}

// NotNull drop not null
func (a *alterTableActionDrop) NotNull() *alterTableActionDrop {
	a.ordered.Add(0, a.ordered.Concat("DROP NOT NULL"))
	return a
}

// Default drop default
func (a *alterTableActionDrop) Default() *alterTableActionDrop {
	a.ordered.Add(0, a.ordered.Concat("DROP DEFAULT"))
	return a
}

// Expression drop expression
func (a *alterTableActionDrop) Expression() *alterTableActionDrop {
	a.ordered.Add(0, a.ordered.Concat("DROP EXPRESSION"))
	return a
}

// Identity drop identity
func (a *alterTableActionDrop) Identity() *alterTableActionDrop {
	a.ordered.Add(0, a.ordered.Concat("DROP IDENTITY"))
	return a
}

// Constraint drop constraint
func (a *alterTableActionDrop) Constraint(name string) *alterTableActionDrop {
	a.ordered.Add(0, a.ordered.Concat("DROP CONSTRAINT"))
	a.ordered.Add(2, a.ordered.Concat(name))
	return a
}

// IsEmpty check if empty
func (a *alterTableActionDrop) IsEmpty() bool {
	return a == nil || a.ordered.IsEmpty()
}

// Reset reset item data
func (a *alterTableActionDrop) Reset() *alterTableActionDrop {
	a.ordered.Reset()
	return a
}

// Grow memory
func (a *alterTableActionDrop) Grow(n int) *alterTableActionDrop {
	a.ordered.Grow(n)
	return a
}

// GetArguments get arguments
func (a *alterTableActionDrop) GetArguments() []any {
	return a.ordered.GetArguments()
}

// String render alter table query
func (a *alterTableActionDrop) String() string {
	if a.IsEmpty() {
		return ""
	}
	return a.ordered.String()
}
