package querymodel

import validation "github.com/go-ozzo/ozzo-validation/v4"

// ExpenseMoneyAll ...
type ExpenseMoneyAll struct {
	PageToken string `query:"pageToken"`
	FromAt    string `query:"fromAt"`
	ToAt      string `query:"toAt"`
	Sort      string `query:"sort"`
}

// Validate ...
func (m ExpenseMoneyAll) Validate() error {
	return validation.ValidateStruct(&m)
}
