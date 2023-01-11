package responsemodel

import "expense-tracker-server/external/util/ptime"

// ResponseCategoryAll ...
type ResponseCategoryAll struct {
	List  []ResponseCategoryAdmin `json:"list"`
	Total int64                   `json:"total"`
	Limit int64                   `json:"limit"`
}

// ResponseCategoryAdmin ...
type ResponseCategoryAdmin struct {
	ID        string              `json:"_id"`
	Name      string              `json:"name"`
	Type      string              `json:"type"`
	Status    string              `json:"status"`
	CreatedAt *ptime.TimeResponse `json:"createdAt"`
	UpdatedAt *ptime.TimeResponse `json:"updatedAt"`
}
