package gosql

import "strings"

// ISQL each query should implement the interface
type ISQL interface {
	// SQL Get query as string with all params
	SQL() (query string, params []any, returning []any)
}

// Check Table for ISQL
var _ = ISQL(&Table{})

// Check Index for ISQL
var _ = ISQL(&Index{})

// Check Comment for ISQL
var _ = ISQL(&Comment{})

// Check Select for ISQL
var _ = ISQL(&Select{})

// Check Insert for ISQL
var _ = ISQL(&Insert{})

// Check Update for ISQL
var _ = ISQL(&Update{})

// Check Delete for ISQL
var _ = ISQL(&Delete{})

// Check Merge for ISQL
var _ = ISQL(&Merge{})

// Check Alter for ISQL
var _ = ISQL(&Alter{})

// SQList Collection of SQL element
type SQList []ISQL

// Join Return query string joined in one string
func (s SQList) Join() (query string, params []any, returning []any) {
	b := strings.Builder{}
	for _, isql := range s {
		q, p, r := isql.SQL()
		if q[len(q)-1] != ';' {
			b.WriteString(q + ";")
		} else {
			b.WriteString(q)
		}
		params = append(params, p...)
		returning = append(returning, r...)
	}
	query = b.String()
	return
}
