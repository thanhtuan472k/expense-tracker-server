package errorcode

import "expense-tracker-server/external/response"

const (
	CategoryIsRequiredName  = "category_is_required_name"
	CategoryIsInvalid       = "category_is_invalid"
	CategoryStatusIsInvalid = "category_status_is_invalid"
	CategoryIDIsInvalid     = "category_id_is_invalid"
	CategoryNotFound        = "category_not_found"
	CategoryTypeIsInvalid   = "category_type_is_invalid"
	CategoryStatusInactive  = "category_status_inactive"
)

// 300 - 399
var category = []response.Code{
	{
		Key:     CategoryIsRequiredName,
		Message: "Danh mục không được trống",
		Code:    300,
	},
	{
		Key:     CategoryIsInvalid,
		Message: "Danh mục không hợp lệ",
		Code:    301,
	},
	{
		Key:     CategoryStatusIsInvalid,
		Message: "Trạng thái danh mục không hợp lệ",
		Code:    302,
	},
	{
		Key:     CategoryIDIsInvalid,
		Message: "ID danh mục không hợp lệ",
		Code:    303,
	},
	{
		Key:     CategoryNotFound,
		Message: "Không tìm thấy danh mục",
		Code:    304,
	},
	{
		Key:     CategoryTypeIsInvalid,
		Message: "Loại danh mục không hợp lệ",
		Code:    305,
	},
	{
		Key:     CategoryStatusInactive,
		Message: "Danh mục đã bị ẩn",
		Code:    306,
	},
}
