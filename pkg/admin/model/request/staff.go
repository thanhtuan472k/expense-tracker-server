package requestmodel

import validation "github.com/go-ozzo/ozzo-validation/v4"

// StaffBodyLogin ...
type StaffBodyLogin struct {
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate ...
func (m StaffBodyLogin) Validate() error {
	return validation.ValidateStruct(&m)
}
