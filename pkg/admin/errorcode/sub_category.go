package errorcode

import "expense-tracker-server/external/response"

const (
	SubCategoryIDIsInvalid = "category_is_required_name"
)

// 400 - 499
var subCategory = []response.Code{
	{
		Key:     SubCategoryIDIsInvalid,
		Message: "Danh mục không được trống",
		Code:    400,
	},
}
