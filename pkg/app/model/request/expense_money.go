package requestmodel

import validation "github.com/go-ozzo/ozzo-validation/v4"

// ExpenseMoneyBodyCreate ...
type ExpenseMoneyBodyCreate struct {
}

// Validate ...
func (m ExpenseMoneyBodyCreate) Validate() error {
	return validation.ValidateStruct(&m)
}

// ExpenseMoneyBodyUpdate ...
type ExpenseMoneyBodyUpdate struct {
}

// Validate ...
func (m ExpenseMoneyBodyUpdate) Validate() error {
	return validation.ValidateStruct(&m)
}
