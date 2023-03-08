package requestmodel

import validation "github.com/go-ozzo/ozzo-validation/v4"

//
// Register
//

// UserBodyRegister ...
type UserBodyRegister struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Validate ...
func (m UserBodyRegister) Validate() error {
	return validation.ValidateStruct(&m)
}

//
// Login
//

// UserBodyLogin ...
type UserBodyLogin struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

// Validate ...
func (m UserBodyLogin) Validate() error {
	return validation.ValidateStruct(&m)
}
