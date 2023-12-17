package gosql

import (
	"strings"
	"time"
	"unicode/utf8"
)

// Sorting fields
// Example: ['createdAt:desc', 'name', 'qty:asc']
type Sorting []string

// Allowed return all sorting columns according to allowed sort map
func (s Sorting) Allowed(items map[string]string) []string {
	var result = make([]string, 0, len(s))
	for _, field := range s {
		parts := strings.Split(strings.Trim(field, " 	"), ":")
		if v, ok := items[parts[0]]; ok {
			if len(parts) > 1 {
				if strings.ToLower(parts[1]) == "desc" {
					v += " DESC"
				}
			}
			result = append(result, v)
		}
	}
	return result
}

// PeriodFilter filter by datetime columns
type PeriodFilter struct {
	// begins from
	Start *time.Time `json:"start"`
	// ended when
	End *time.Time `json:"end"`
}

// IsEmpty if filter is empty
func (p *PeriodFilter) IsEmpty() bool {
	return p == nil || p.Start == nil && p.End == nil
}

// FieldCondition получить условие для фильтрации
func (p *PeriodFilter) FieldCondition(field string) *Condition {
	cond := NewSqlCondition(ConditionOperatorAnd)
	if p.Start != nil {
		cond.AddExpression(field+" >= ?", p.Start.Local())
	}
	if p.End != nil {
		cond.AddExpression(field+" <= ?", p.End.Local())
	}
	return cond
}

// SearchString search by text columns
type SearchString string

// PrepareLikeValue prepare search like condition
func (s SearchString) PrepareLikeValue(column string) *Condition {
	cond := NewSqlCondition(ConditionOperatorAnd)
	normal := strings.ToLower(strings.Trim(string(s), " 	"))
	if utf8.RuneCountInString(normal) > 0 {
		parts := strings.Split(normal, " ")
		for _, part := range parts {
			cond.AddExpression(column+" like lower(?)", "%"+part+"%")
		}
	}
	return cond
}
