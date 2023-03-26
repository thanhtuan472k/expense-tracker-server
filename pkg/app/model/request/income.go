package requestmodel

import validation "github.com/go-ozzo/ozzo-validation/v4"

// IncomeBodyCreate ...
type IncomeBodyCreate struct {
	Name        string  `json:"name"`
	Money       float64 `json:"money"`
	Category    string  `json:"category"`
	Description string  `json:"description"`
}

// Validate ...
func (m IncomeBodyCreate) Validate() error {
	return validation.ValidateStruct(&m)
}
