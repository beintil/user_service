package client

import (
	"github.com/dimonrus/gohelp"
	"time"
)

// Person entity stores information about a person
type Person struct {
	// Identification of the person
	Id *int64 `db:"id" json:"id"`
	// Identification of the user
	UserId *int64 `db:"user_id" json:"userId"`
	// First name
	FirstName *string `db:"first_name" json:"firstName" validate:"required"`
	// Last name
	LastName *string `db:"last_name" json:"lastName" validate:"required"`
	// Middle name
	MiddleName *string `db:"middle_name" json:"middleName,omitempty"`
	// Birth date
	Birthday *time.Time `db:"birthday" json:"birthday,omitempty"`
	// Age
	Age *uint8 `db:"age" json:"age,omitempty"`
	// Gender
	Gender *string `db:"gender" json:"gender,omitempty" validate:"omitempty,oneof=M F"`
	// Phone number
	Phone *string `db:"phone_number" json:"phone,omitempty"`
	// CreatedAt
	CreatedAt *time.Time `db:"created_at" json:"createdAt"`
	// UpdatedAt
	UpdatedAt *time.Time `db:"updated_at" json:"updated,omitempty"`
}

func (m *Person) Table() string {
	return `"person"`
}

func (m *Person) Values() []interface{} {
	return []interface{}{&m.Id, &m.UserId, &m.FirstName, &m.LastName, &m.MiddleName, &m.Birthday, &m.Age, &m.Gender, &m.Phone, &m.CreatedAt, &m.UpdatedAt}
}

func (m *Person) Columns() []string {
	return []string{"id", "user_id", "first_name", "last_name", "middle_name", "birthday", "age", "gender", "phone_number", "created_at", "updated_at"}
}

func (m *Person) SetCreatedAt() {
	m.CreatedAt = gohelp.Ptr(time.Now())
}

func (m *Person) SetUpdateAt() {
	m.UpdatedAt = gohelp.Ptr(time.Now())
}
