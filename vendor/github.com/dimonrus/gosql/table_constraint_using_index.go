package gosql

// [ CONSTRAINT constraint_name ]
// { UNIQUE | PRIMARY KEY } USING INDEX index_name
// [ DEFERRABLE | NOT DEFERRABLE ] [ INITIALLY DEFERRED | INITIALLY IMMEDIATE ]
type constraintTableUsingIndex struct {
	// ordered expression
	ordered orderedExpression
}

// Name set constraint name
func (c *constraintTableUsingIndex) Name(name string) *constraintTableUsingIndex {
	c.ordered.Add(0, c.ordered.Concat("CONSTRAINT ", name))
	return c
}

// Unique set unique
func (c *constraintTableUsingIndex) Unique() *constraintTableUsingIndex {
	c.ordered.Add(1, c.ordered.Concat("UNIQUE"))
	return c
}

// PrimaryKey set primary key
func (c *constraintTableUsingIndex) PrimaryKey() *constraintTableUsingIndex {
	c.ordered.Add(1, c.ordered.Concat("PRIMARY KEY"))
	return c
}

// Using set using index name
func (c *constraintTableUsingIndex) Using(name string) *constraintTableUsingIndex {
	c.ordered.Add(2, c.ordered.Concat("USING INDEX ", name))
	return c
}

// Deferrable set deferrable
func (c *constraintTableUsingIndex) Deferrable() *constraintTableUsingIndex {
	c.ordered.Add(3, c.ordered.Concat("DEFERRABLE"))
	return c
}

// NotDeferrable set not deferrable
func (c *constraintTableUsingIndex) NotDeferrable() *constraintTableUsingIndex {
	c.ordered.Add(3, c.ordered.Concat("NOT DEFERRABLE"))
	return c
}

// InitiallyDeferred set INITIALLY DEFERRED
func (c *constraintTableUsingIndex) InitiallyDeferred() *constraintTableUsingIndex {
	c.ordered.Add(4, c.ordered.Concat("INITIALLY DEFERRED"))
	return c
}

// InitiallyImmediate set INITIALLY IMMEDIATE
func (c *constraintTableUsingIndex) InitiallyImmediate() *constraintTableUsingIndex {
	c.ordered.Add(4, c.ordered.Concat("INITIALLY IMMEDIATE"))
	return c
}

// IsEmpty check if empty
func (c *constraintTableUsingIndex) IsEmpty() bool {
	return c == nil || c.ordered.IsEmpty()
}

// Reset reset item data
func (c *constraintTableUsingIndex) Reset() *constraintTableUsingIndex {
	c.ordered.Reset()
	return c
}

// Grow memory
func (c *constraintTableUsingIndex) Grow(n int) *constraintTableUsingIndex {
	c.ordered.Grow(n)
	return c
}

// GetArguments get arguments
func (c *constraintTableUsingIndex) GetArguments() []any {
	return c.ordered.GetArguments()
}

// String render alter table query
func (c *constraintTableUsingIndex) String() string {
	if c.IsEmpty() {
		return ""
	}
	return c.ordered.String()
}
