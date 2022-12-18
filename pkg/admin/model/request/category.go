package requestmodel

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

// CategoryBody ...
type CategoryBody struct {
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Validate ...
func (b CategoryBody) Validate() error {
	return validation.ValidateStruct(&b)
}
