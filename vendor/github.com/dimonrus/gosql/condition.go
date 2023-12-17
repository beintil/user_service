package gosql

import "strings"

const (
	ConditionOperatorAnd = "AND"
	ConditionOperatorOr  = "OR"
	ConditionOperatorXor = "XOR"
)

// Merge with Condition
type merge struct {
	operator  string
	condition []*Condition
}

// Condition type
type Condition struct {
	operator   string
	expression []string
	argument   []interface{}
	merge      *merge
}

// NewSqlCondition init condition
func NewSqlCondition(operator string) *Condition {
	return &Condition{operator: operator, expression: make([]string, 0, 8), argument: make([]interface{}, 0, 8)}
}

// Get string of conditions
func (c *Condition) String() string {
	if c.merge != nil {
		var slaves []string
		for i := range (*c.merge).condition {
			slaves = append(slaves, (*c.merge).condition[i].String())
		}
		if c.expression != nil {
			slaves = append(slaves, "("+strings.Join(c.expression, " "+c.operator+" ")+")")
		}
		return "(" + strings.Join(slaves, " "+c.merge.operator+" ") + ")"
	} else {
		if c.expression != nil {
			return "(" + strings.Join(c.expression, " "+c.operator+" ") + ")"
		}
		return ""
	}
}

// IsEmpty check if condition is empty
func (c *Condition) IsEmpty() bool {
	return c == nil || (len(c.expression) == 0 && c.merge == nil)
}

// GetArguments get arguments
func (c *Condition) GetArguments() []interface{} {
	var arguments = make([]interface{}, 0, 4)
	if c.merge != nil {
		for i := range (*c.merge).condition {
			arguments = append(arguments, (*c.merge).condition[i].GetArguments()...)
		}
	}
	return append(arguments, c.argument...)
}

// AddExpression add expression
func (c *Condition) AddExpression(expression string, values ...interface{}) *Condition {
	c.expression = append(c.expression, expression)
	c.argument = append(c.argument, values...)
	return c
}

// AddArgument add argument
func (c *Condition) AddArgument(values ...interface{}) *Condition {
	c.argument = append(c.argument, values...)
	return c
}

// Replace current condition
func (c *Condition) Replace(cond *Condition) *Condition {
	*c = *cond
	return c
}

// Merge with conditions
func (c *Condition) Merge(operator string, conditions ...*Condition) *Condition {
	for i := range conditions {
		if conditions[i].IsEmpty() {
			continue
		}
		if c.merge == nil {
			c.merge = &merge{operator: operator, condition: make([]*Condition, 0, len(conditions))}
		}
		c.merge.condition = append(c.merge.condition, conditions[i])
	}
	return c
}
