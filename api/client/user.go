package client

import (
	"github.com/dimonrus/gohelp"
	"time"
)

// User entity for authentication
type User struct {
	// Id user identification
	Id *int64 `db:"id" json:"id"`
	// Password
	Password *string `db:"password" json:"password,omitempty"  validate:"omitempty,min=5,max=20" validate:"required_if=Password"`
	// NewPassword New password
	NewPassword *string `json:"newPassword,omitempty" validate:"omitempty,min=5,max=20" validate:"required_if=Password"`
	// Email
	Email *string `db:"email" json:"email" validate:"required,email"`
	// CreatedAt
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	// UpdatedAt
	UpdatedAt *time.Time `db:"updated_at" json:"updated,omitempty"`
	// Person stores info about the user
	Person Person `db:"person" json:"person"`
}

func (m *User) Table() string {
	return `"user"`
}

func (m *User) ClearSensitiveData() *User {
	m.Password = nil
	m.NewPassword = nil
	return m
}

func (m *User) Values() []interface{} {
	return []any{&m.Id, &m.Password, &m.Email, &m.CreatedAt, &m.UpdatedAt}
}
func (m *User) ValuesWithPerson() []interface{} {
	var values = make([]interface{}, 0, len(m.Person.Values())+len(m.Values()))
	values = append(values, &m.Id, &m.Email, &m.CreatedAt, &m.UpdatedAt)
	return append(values, m.Person.Values()...)
}

func (m *User) Columns() []string {
	return []string{"id", "password", "email", "created_at", "updated_at"}
}

func (m *User) SetCreatedAt() {
	m.CreatedAt = gohelp.Ptr(time.Now())
}

func (m *User) SetUpdateAt() {
	m.UpdatedAt = gohelp.Ptr(time.Now())
}

func NewUser() *User {
	return &User{Person: Person{}}
}
