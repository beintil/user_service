package client

import (
	"github.com/dimonrus/gosql"
	"github.com/lib/pq"
)

// Users lists user
type Users []*User

// UserSearchForm user search form
type UserSearchForm struct {
	// Identification of the users
	Ids []int32 `json:"ids"`
	// Name of the user
	LastNames []string `json:"lastNames"`
	// First name of the user
	FirstNames []string `json:"firstNames"`
	// Middle name of the user
	MiddleNames []string `json:"middleNames"`
}

func (f *UserSearchForm) PrepareCondition() *gosql.Condition {
	cond := gosql.NewSqlCondition(gosql.ConditionOperatorAnd)
	if len(f.Ids) > 0 {
		cond.AddExpression("u.id = ANY(?)", pq.Array(f.Ids))
	}
	if len(f.LastNames) > 0 {
		cond.AddExpression("p.last_name = ANY(?)", pq.Array(f.LastNames))
	}
	if len(f.FirstNames) > 0 {
		cond.AddExpression("p.first_name = ANY(?)", pq.Array(f.FirstNames))
	}
	if len(f.MiddleNames) > 0 {
		cond.AddExpression("p.middle_name = ANY(?)", pq.Array(f.MiddleNames))
	}
	return cond
}
