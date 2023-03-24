package querymodel

import validation "github.com/go-ozzo/ozzo-validation/v4"

// CategoryAll ...
type CategoryAll struct {
	PageToken string `json:"pageToken" query:"pageToken"`
	Status    string `json:"status" query:"status"`
	Sort      string `json:"sort" query:"sort"`
}

// Validate ...
func (m CategoryAll) Validate() error {
	return validation.ValidateStruct(&m)
}
