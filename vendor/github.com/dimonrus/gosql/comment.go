package gosql

import "strings"

// COMMENT ON
//{
//  ACCESS METHOD object_name |
//  AGGREGATE aggregate_name ( aggregate_signature ) |
//  CAST (source_type AS target_type) |
//  COLLATION object_name |
//  COLUMN relation_name.column_name |
//  CONSTRAINT constraint_name ON table_name |
//  CONSTRAINT constraint_name ON DOMAIN domain_name |
//  CONVERSION object_name |
//  DATABASE object_name |
//  DOMAIN object_name |
//  EXTENSION object_name |
//  EVENT TRIGGER object_name |
//  FOREIGN DATA WRAPPER object_name |
//  FOREIGN TABLE object_name |
//  FUNCTION function_name [ ( [ [ argmode ] [ argname ] argtype [, ...] ] ) ] |
//  INDEX object_name |
//  LARGE OBJECT large_object_oid |
//  MATERIALIZED VIEW object_name |
//  OPERATOR operator_name (left_type, right_type) |
//  OPERATOR CLASS object_name USING index_method |
//  OPERATOR FAMILY object_name USING index_method |
//  POLICY policy_name ON table_name |
//  [ PROCEDURAL ] LANGUAGE object_name |
//  PROCEDURE procedure_name [ ( [ [ argmode ] [ argname ] argtype [, ...] ] ) ] |
//  PUBLICATION object_name |
//  ROLE object_name |
//  ROUTINE routine_name [ ( [ [ argmode ] [ argname ] argtype [, ...] ] ) ] |
//  RULE rule_name ON table_name |
//  SCHEMA object_name |
//  SEQUENCE object_name |
//  SERVER object_name |
//  STATISTICS object_name |
//  SUBSCRIPTION object_name |
//  TABLE object_name |
//  TABLESPACE object_name |
//  TEXT SEARCH CONFIGURATION object_name |
//  TEXT SEARCH DICTIONARY object_name |
//  TEXT SEARCH PARSER object_name |
//  TEXT SEARCH TEMPLATE object_name |
//  TRANSFORM FOR type_name LANGUAGE lang_name |
//  TRIGGER trigger_name ON table_name |
//  TYPE object_name |
//  VIEW object_name
// } IS 'text'

type Comment struct {
	detailedExpression detailedExpression
}

// IsEmpty check if empty
func (c *Comment) IsEmpty() bool {
	return c == nil || c.detailedExpression.IsEmpty()
}

// Column comment column
func (c *Comment) Column(column string, comment string) *Comment {
	c.detailedExpression.SetDetail("COLUMN " + column)
	c.detailedExpression.Expression().Add(comment)
	return c
}

// Table comment table
func (c *Comment) Table(table string, comment string) *Comment {
	c.detailedExpression.SetDetail("TABLE " + table)
	c.detailedExpression.Expression().Add(comment)
	return c
}

// String render comment query
func (c *Comment) String() string {
	if c.IsEmpty() {
		return ""
	}
	b := strings.Builder{}
	b.WriteString("COMMENT ON ")
	if c.detailedExpression.GetDetail() != "" {
		b.WriteString(c.detailedExpression.GetDetail() + " IS '" + c.detailedExpression.Expression().String(EnumDelimiter) + "';")
	}
	return b.String()
}

// SQL common sql interface
func (c *Comment) SQL() (query string, params []any, returning []any) {
	query = c.String()
	return
}

// NewComment init comment
func NewComment() *Comment {
	return &Comment{}
}
