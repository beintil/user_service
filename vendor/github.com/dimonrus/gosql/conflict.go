package gosql

import "strings"

// On conflict query part
type conflict struct {
	// object of conflict
	object string
	// action on conflict
	action string
	// set of changes
	set expression
	// condition
	where Condition
	// constraint
	constraint string
}

// String conflict expression
func (c *conflict) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("ON CONFLICT")
	if c.object != "" {
		b.WriteString(" (" + c.object + ")")
	}
	if c.constraint != "" {
		b.WriteString(" ON CONSTRAINT " + c.constraint)
	}
	if c.set.Len() > 0 {
		if c.action != "" {
			b.WriteString(" DO " + c.action)
		}
		b.WriteString(" SET " + c.set.String(", "))
		if !c.where.IsEmpty() {
			b.WriteString(" WHERE " + c.where.String())
		}
	} else {
		if !c.where.IsEmpty() {
			b.WriteString(" WHERE " + c.where.String())
		}
		if c.action != "" {
			b.WriteString(" DO " + c.action)
		}
	}
	return b.String()
}

// GetArguments get all arguments
func (c *conflict) GetArguments() []any {
	return append(c.set.GetArguments(), c.where.GetArguments()...)
}

// IsEmpty Is conflict empty
func (c *conflict) IsEmpty() bool {
	return c.object == "" && c.action == "" && c.set.Len() == 0 && c.where.IsEmpty() && c.constraint == ""
}

// Object of conflict
func (c *conflict) Object(object string) *conflict {
	c.object = object
	return c
}

// GetObject get conflict object
func (c *conflict) GetObject() string {
	return c.object
}

// ResetObject reset
func (c *conflict) ResetObject() *conflict {
	c.object = ""
	return c
}

// Action of conflict
func (c *conflict) Action(action string) *conflict {
	c.action = action
	return c
}

// GetAction get action
func (c *conflict) GetAction() string {
	return c.action
}

// ResetAction reset action
func (c *conflict) ResetAction() *conflict {
	c.action = ""
	return c
}

// Set of expressions on conflict
func (c *conflict) Set() *expression {
	return &c.set
}

// Where get condition
func (c *conflict) Where() *Condition {
	return &c.where
}

// Constraint set constraint
func (c *conflict) Constraint(constraint string) *conflict {
	c.constraint = constraint
	return c
}

// GetConstraint get constraint
func (c *conflict) GetConstraint() string {
	return c.constraint
}

// ResetConstraint reset constraint
func (c *conflict) ResetConstraint() *conflict {
	c.constraint = ""
	return c
}

// NewConflict conflict constructor
func NewConflict() *conflict {
	return &conflict{
		where: Condition{operator: ConditionOperatorAnd},
	}
}
