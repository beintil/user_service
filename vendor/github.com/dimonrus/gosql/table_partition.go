package gosql

import "strings"

// { column_name | ( expression ) } [ COLLATE collation ] [ opclass ]
type partitionColumn struct {
	expression expression
	// collation
	collate string
	// opclass
	opclass string
}

// Expression get expression
func (p *partitionColumn) Expression() *expression {
	return &p.expression
}

// Collate set sort rule
func (p *partitionColumn) Collate(collation string) *partitionColumn {
	p.collate = collation
	return p
}

// OpClass set opclass
func (p *partitionColumn) OpClass(opclass string) *partitionColumn {
	p.opclass = opclass
	return p
}

// IsEmpty check is partition is empty
func (p *partitionColumn) IsEmpty() bool {
	return p == nil || (p.expression.Len() == 0 && p.collate == "" && p.opclass == "")
}

// String Render partition into string
func (p *partitionColumn) String() string {
	if p.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if p.expression.Len() > 0 {
		b.WriteString(" (" + p.expression.String(", ") + ")")
	}
	if p.collate != "" {
		b.WriteString(" COLLATE " + p.collate)
	}
	if p.opclass != "" {
		b.WriteString(" " + p.opclass)
	}
	return b.String()
}

// [ PARTITION BY { RANGE | LIST | HASH } ( { column_name | ( expression ) } [ COLLATE collation ] [ opclass ] [, ... ] ) ]
type partitionTable struct {
	// type of partition
	partitionBy string
	// partition columns list
	clause expression
}

// By type of partition RANGE | LIST | HASH
func (p *partitionTable) By(class string) *partitionTable {
	p.partitionBy = class
	return p
}

// Clause Add partition clause
// { column_name | ( expression ) } [ COLLATE collation ] [ opclass ] [, ... ]
func (p *partitionTable) Clause(expression ...string) *expression {
	p.clause.Add(expression...)
	return &p.clause
}

// IsEmpty check is partition is empty
func (p *partitionTable) IsEmpty() bool {
	return p == nil || (p.partitionBy == "" && p.clause.Len() == 0)
}

// String Render partition into string
func (p *partitionTable) String() string {
	if p.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if p.partitionBy != "" {
		b.WriteString(" " + p.partitionBy)
	}
	if p.clause.Len() > 0 {
		b.WriteString(" (" + p.clause.String(", ") + ")")
	}
	return b.String()[1:]
}

// partition_bound_spec
// IN ( partition_bound_expr [, ...] ) |
// FROM ( { partition_bound_expr | MINVALUE | MAXVALUE } [, ...] )
//
//	TO ( { partition_bound_expr | MINVALUE | MAXVALUE } [, ...] ) |
//
// WITH ( MODULUS numeric_literal, REMAINDER numeric_literal )
type partitionBound struct {
	// in
	in expression
	// from
	from expression
	// to
	to expression
	// with
	with expression
}

// In expression
func (p *partitionBound) In() *expression {
	return &p.in
}

// From expression
func (p *partitionBound) From() *expression {
	return &p.from
}

// To expression
func (p *partitionBound) To() *expression {
	return &p.to
}

// With expression
func (p *partitionBound) With() *expression {
	return &p.with
}

// IsEmpty check is partition bound is empty
func (p *partitionBound) IsEmpty() bool {
	return p == nil || (p.in.Len() == 0 && p.from.Len() == 0 && p.to.Len() == 0 && p.with.Len() == 0)
}

// String Render partition bound into string
func (p *partitionBound) String() string {
	if p.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	if p.in.Len() > 0 {
		b.WriteString(" IN (" + p.in.String(", ") + ")")
	}
	if p.from.Len() > 0 {
		b.WriteString(" FROM (" + p.from.String(", ") + ")")
	}
	if p.to.Len() > 0 {
		b.WriteString(" TO (" + p.to.String(", ") + ")")
	}
	if p.with.Len() > 0 {
		b.WriteString(" WITH (" + p.with.String(", ") + ")")
	}

	return b.String()[1:]
}
