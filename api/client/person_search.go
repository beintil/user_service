package client

import (
	"github.com/dimonrus/gosql"
	"github.com/lib/pq"
)

// Persons list Person
type Persons []*Person

// PersonSearchForm search form
type PersonSearchForm struct {
	// Identification of the peron
	Ids []int32 `json:"ids"`
	// Identification of the users
	UserIds []int32 `json:"userIds"`
	// Last name of the person
	LastNames []string `json:"lastNames"`
	// First name of the person
	FirstNames []string `json:"firstNames"`
	// Middle name of the person
	MiddleNames []string `json:"middleNames"`
}

func (f *PersonSearchForm) PrepareCondition() *gosql.Condition {
	cond := gosql.NewSqlCondition(gosql.ConditionOperatorAnd)
	if len(f.Ids) > 0 {
		cond.AddExpression("id = ANY(?)", pq.Array(f.Ids))
	}
	if len(f.UserIds) > 0 {
		cond.AddExpression("user_id = ANY(?)", pq.Array(f.UserIds))
	}
	if len(f.LastNames) > 0 {
		cond.AddExpression("last_name = ANY(?)", pq.Array(f.LastNames))
	}
	if len(f.FirstNames) > 0 {
		cond.AddExpression("first_name = ANY(?)", pq.Array(f.FirstNames))
	}
	if len(f.MiddleNames) > 0 {
		cond.AddExpression("middle_name = ANY(?)", pq.Array(f.MiddleNames))
	}
	return cond
}
