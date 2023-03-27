package querymodel

// IncomeMoneyAll ...
type IncomeMoneyAll struct {
	PageToken string `query:"pageToken"`
	FromAt    string `query:"fromAt"`
	ToAt      string `query:"toAt"`
	Sort      string `query:"sort"`
}
