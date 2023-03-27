package errorcode

import "expense-tracker-server/external/response"

const (
	CategoryNotFound  = "user_not_found"
	CategoryIsInvalid = "category_is_invalid"
)

// 300 - 399
var category = []response.Code{
	{
		Key:     CategoryNotFound,
		Message: "Danh mục không tồn tại",
		Code:    300,
	},
	{
		Key:     CategoryIsInvalid,
		Message: "Danh mục không hợp lệ",
		Code:    301,
	},
}
