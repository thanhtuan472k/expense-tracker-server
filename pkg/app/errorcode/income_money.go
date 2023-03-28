package errorcode

import "expense-tracker-server/external/response"

const (
	IncomeMoneyNotFound = "income_money_not_found"
)

// 400 - 499
var incomeMoney = []response.Code{
	{
		Key:     IncomeMoneyNotFound,
		Message: "Thu nhập không tìm thấy",
		Code:    400,
	},
}
