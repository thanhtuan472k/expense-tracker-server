package errorcode

import "expense-tracker-server/external/response"

const (
	ExpenseMoneyNotFound = "income_money_not_found"
)

// 600 - 699
var expenseMoney = []response.Code{
	{
		Key:     ExpenseMoneyNotFound,
		Message: "Chi tiêu không tìm thấy",
		Code:    600,
	},
}
