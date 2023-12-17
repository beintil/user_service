package gohelp

import (
	"fmt"
	"regexp"
)

var matchUUIDPattern = regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$")

// UUID type
type UUID struct {
	value *string
}

// NewUUID uuid constructor
func NewUUID() string {
	return *(&UUID{}).Generate().Get()
}

// Generate uuid
func (u *UUID) Generate() *UUID {
	b := make([]byte, 16)
	_, err := rnd.Read(b)
	if err != nil {
		return nil
	}
	uuid := fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	u.value = &uuid
	return u
}

// Reset uuid
func (u *UUID) Reset() *UUID {
	u.value = nil
	return u
}

// Get uuid
func (u *UUID) Get() *string {
	return u.value
}

// IsValid is uuid valid
func (u *UUID) IsValid() bool {
	return matchUUIDPattern.MatchString(*u.Get())
}

// ParseUUID uuid parser
func ParseUUID(data string) *UUID {
	uuid := &UUID{value: &data}
	if uuid.IsValid() {
		return uuid
	}
	return nil
}
