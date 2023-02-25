package errorcode

import "expense-tracker-server/external/response"

const (
	SubCategoryIsRequiredName = "category_is_required_name"
	SubCategoryNotFound       = "sub_category_not_found"
)

// 400 - 499
var subCategory = []response.Code{
	{
		Key:     SubCategoryIsRequiredName,
		Message: "Tên danh mục không được trống",
		Code:    400,
	},
	{
		Key:     SubCategoryNotFound,
		Message: "Danh mục con không tìm thấy",
		Code:    401,
	},
}
