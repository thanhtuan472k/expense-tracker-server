package responsemodel

import (
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/util/ptime"
)

// ResponseExpenseMoneyAll ...
type ResponseExpenseMoneyAll struct {
	List          []ResponseExpenseMoneyInfo `json:"list"`
	EndData       bool                       `json:"endData"`
	NextPageToken string                     `json:"nextPageToken"`
	Total         int64                      `json:"total"`
}

// ResponseExpenseMoneyInfo ...
type ResponseExpenseMoneyInfo struct {
	ID          string                  `json:"_id"`
	Category    mgexpense.CategoryShort `json:"category"`
	SubCategory mgexpense.SubCategory   `json:"subCategory"`
	Money       float64                 `json:"money"`
	Note        string                  `json:"note"`
	CreatedAt   *ptime.TimeResponse     `json:"createdAt"`
	UpdatedAt   *ptime.TimeResponse     `json:"updatedAt"`
}
