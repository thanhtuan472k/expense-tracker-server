package errorcode

import "expense-tracker-server/external/response"

const (
	CategoryIsRequiredName = "category_is_required_name"
	CategoryIsInvalid      = "category_is_invalid"
)

// 200 - 299
var category = []response.Code{
	{
		Key:     CategoryIsRequiredName,
		Message: "Danh mục không được trống",
		Code:    200,
	},
	{
		Key:     CategoryIsInvalid,
		Message: "Danh mục không hợp lệ",
		Code:    201,
	},
}
