package responsemodel

import (
	mgexpense "expense-tracker-server/external/model/mg/expense"
	"expense-tracker-server/external/util/ptime"
)

// ResponseIncomeMoneyAll ...
type ResponseIncomeMoneyAll struct {
	List          []ResponseIncomeMoneyInfo `json:"list"`
	EndData       bool                      `json:"endData"`
	NextPageToken string                    `json:"nextPageToken"`
}

// ResponseIncomeMoneyInfo ...
type ResponseIncomeMoneyInfo struct { // meta struct for api list and detail
	ID        string                  `json:"_id"`
	Category  mgexpense.CategoryShort `json:"category"`
	Money     float64                 `json:"money"`
	Note      string                  `json:"note"`
	CreatedAt *ptime.TimeResponse     `json:"createdAt"`
	UpdatedAt *ptime.TimeResponse     `json:"updatedAt"`
}
