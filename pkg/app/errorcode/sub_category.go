package errorcode

import "expense-tracker-server/external/response"

const (
	SubCategoryNotFound  = "sub_category_not_found"
	SubCategoryIsInvalid = "sub_category_is_invalid"
)

// 300 - 399
var subCategory = []response.Code{
	{
		Key:     SubCategoryNotFound,
		Message: "Danh mục con không tồn tại",
		Code:    300,
	},
	{
		Key:     SubCategoryIsInvalid,
		Message: "Danh mục con không hợp lệ",
		Code:    301,
	},
}
