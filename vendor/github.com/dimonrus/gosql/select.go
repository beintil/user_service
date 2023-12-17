package gosql

import (
	"fmt"
	"strings"
)

// SQL Pagination limit offset
type pagination struct {
	// limit
	Limit int
	// offset
	Offset int
}

// Select Query Builder struct
// Not a thread safety
type Select struct {
	// with queries
	with with
	// columns
	columns expression
	// form rows
	from []string
	// join relations
	join []string
	// where condition
	where Condition
	// order expressions
	orders []string
	// group by expression
	group []string
	// union expressions
	union []*Select
	// except expressions
	except []*Select
	// intersect expressions
	intersect []*Select
	// having conditions
	having Condition
	// pagination for pages
	pagination pagination
	// is subquery. Put query in bracers
	SubQuery bool
}

// SQL Get sql query
func (q *Select) SQL() (query string, params []any, returning []any) {
	return q.String(), q.GetArguments(), q.Columns().GetArguments()
}

// With get with queries
func (q *Select) With() *with {
	return &q.with
}

// Union add union
func (q *Select) Union(s *Select) *Select {
	if s != nil {
		q.union = append(q.union, s)
	}
	return q
}

// Except query
func (q *Select) Except(s *Select) *Select {
	if s != nil {
		q.except = append(q.except, s)
	}
	return q
}

// Intersect query
func (q *Select) Intersect(s *Select) *Select {
	if s != nil {
		q.intersect = append(q.intersect, s)
	}
	return q
}

// ResetIntersect reset intersect
func (q *Select) ResetIntersect() *Select {
	q.intersect = make([]*Select, 0)
	return q
}

// ResetUnion reset union
func (q *Select) ResetUnion() *Select {
	q.union = make([]*Select, 0)
	return q
}

// ResetExcept reset except
func (q *Select) ResetExcept() *Select {
	q.except = make([]*Select, 0)
	return q
}

// Append column
func (q *Select) Columns() *expression {
	return &q.columns
}

// Append from
func (q *Select) From(table ...string) *Select {
	q.from = append(q.from, table...)
	return q
}

// Reset column
func (q *Select) ResetFrom() *Select {
	q.from = []string{}
	return q
}

// Append join
func (q *Select) Relate(relation ...string) *Select {
	q.join = append(q.join, relation...)
	return q
}

// Reset join
func (q *Select) ResetRelations() *Select {
	q.join = []string{}
	return q
}

// Where conditions
func (q *Select) Where() *Condition {
	return &q.where
}

// Where conditions
func (q *Select) Having() *Condition {
	return &q.having
}

// Append Order
func (q *Select) AddOrder(expression ...string) *Select {
	q.orders = append(q.orders, expression...)
	return q
}

// Reset Order
func (q *Select) ResetOrder() *Select {
	q.orders = []string{}
	return q
}

// Append Group
func (q *Select) GroupBy(fields ...string) *Select {
	q.group = append(q.group, fields...)
	return q
}

// Reset Group
func (q *Select) ResetGroupBy() *Select {
	q.group = []string{}
	return q
}

// Set pagination
func (q *Select) SetPagination(limit int, offset int) *Select {
	q.pagination = pagination{Limit: limit, Offset: offset}
	return q
}

// Get arguments
func (q *Select) GetArguments() []interface{} {
	arguments := make([]interface{}, 0)
	if q.with.Len() > 0 {
		for _, w := range q.with.queries {
			arguments = append(arguments, w.GetArguments()...)
		}
	}

	arguments = append(arguments, append(q.where.GetArguments(), q.having.GetArguments()...)...)

	if len(q.union) > 0 {
		for _, u := range q.union {
			arguments = append(arguments, u.GetArguments()...)
		}
	}

	if len(q.except) > 0 {
		for _, u := range q.except {
			arguments = append(arguments, u.GetArguments()...)
		}
	}

	if len(q.intersect) > 0 {
		for _, i := range q.intersect {
			arguments = append(arguments, i.GetArguments()...)
		}
	}
	return arguments
}

// Make SQL query
func (q *Select) String() string {
	b := strings.Builder{}

	// With render
	if q.with.Len() > 0 {
		b.WriteString(q.with.String() + " ")
	}

	// Select columns
	if q.columns.Len() > 0 {
		b.WriteString("SELECT " + q.columns.String(", "))
	}

	// From table
	if len(q.from) > 0 {
		b.WriteString(" FROM " + strings.Join(q.from, ", "))
	}

	// From table
	if len(q.join) > 0 {
		b.WriteString(" " + strings.Join(q.join, " "))
	}

	// Where conditions
	if len(q.where.expression) > 0 || q.where.merge != nil {
		b.WriteString(" WHERE " + q.where.String())
	}

	// Prepare groups
	if len(q.group) > 0 {
		b.WriteString(" GROUP BY " + strings.Join(q.group, ", "))
	}

	// Prepare having expression
	if len(q.having.expression) > 0 || q.having.merge != nil {
		b.WriteString(" HAVING " + q.having.String())
	}

	// Prepare orders
	if len(q.orders) > 0 {
		b.WriteString(" ORDER BY " + strings.Join(q.orders, ", "))
	}

	// Prepare pagination
	if q.pagination.Limit > 0 {
		b.WriteString(fmt.Sprintf(" LIMIT %v OFFSET %v", q.pagination.Limit, q.pagination.Offset))
	}

	// Union render
	for _, u := range q.union {
		b.WriteString(" UNION " + u.String())
	}

	// Except render
	for _, u := range q.except {
		b.WriteString(" EXCEPT " + u.String())
	}

	// Intersect render
	for _, i := range q.intersect {
		b.WriteString(" INTERSECT " + i.String())
	}

	// Check if the query is for sub query
	if q.SubQuery {
		return "(" + b.String() + ")"
	}

	return b.String()
}

// NewSelect Query Builder
func NewSelect() *Select {
	return &Select{
		where:  Condition{operator: ConditionOperatorAnd},
		having: Condition{operator: ConditionOperatorAnd},
	}
}
