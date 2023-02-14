package requestmodel

import validation "github.com/go-ozzo/ozzo-validation/v4"

// StaffBodyLogin ...
type StaffBodyLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Validate ...
func (m StaffBodyLogin) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Phone),
	)
}

// StaffBodyUpdate ...
type StaffBodyUpdate struct {
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Gender string `json:"gender"`
}

// Validate ...
func (m StaffBodyUpdate) Validate() error {
	return validation.ValidateStruct(&m)
}

// StaffBodyChangePassword ...
type StaffBodyChangePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// Validate ...
func (m StaffBodyChangePassword) Validate() error {
	return validation.ValidateStruct(&m)
}
