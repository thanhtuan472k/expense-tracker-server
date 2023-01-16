package errorcode

import "expense-tracker-server/external/response"

const (
	CategoryIsRequiredName  = "category_is_required_name"
	CategoryIsInvalid       = "category_is_invalid"
	CategoryStatusIsInvalid = "category_status_is_invalid"
	CategoryIDIsInvalid     = "category_id_is_invalid"
	CategoryNotFound        = "category_not_found"
	CategoryTypeIsInvalid   = "category_type_is_invalid"
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
	{
		Key:     CategoryStatusIsInvalid,
		Message: "Trạng thái danh mục không hợp lệ",
		Code:    202,
	},
	{
		Key:     CategoryIDIsInvalid,
		Message: "ID danh mục không hợp lệ",
		Code:    203,
	},
	{
		Key:     CategoryNotFound,
		Message: "Không tìm thấy danh mục",
		Code:    204,
	},
	{
		Key:     CategoryTypeIsInvalid,
		Message: "Loại danh mục không hợp lệ",
		Code:    205,
	},
}
