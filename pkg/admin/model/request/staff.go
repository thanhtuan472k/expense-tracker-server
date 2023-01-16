package requestmodel

import validation "github.com/go-ozzo/ozzo-validation/v4"

// StaffBodyCreate ...
type StaffBodyCreate struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
	Phone  string `json:"phone"`
}

// Validate ..
func (m StaffBodyCreate) Validate() error {
	return validation.ValidateStruct(&m)
}
