package errorcode

import "expense-tracker-server/external/response"

const (
	StaffNotFound          = "staff_not_found"
	StaffPasswordIncorrect = "staff_password_incorrect"
	StaffStatusInactive    = "staff_status_inactive"
)

// 200 - 299
var staff = []response.Code{
	{
		Key:     StaffNotFound,
		Message: "Không tìm thấy tài khoản admin ",
		Code:    200,
	},
	{
		Key:     StaffPasswordIncorrect,
		Message: "Mật khẩu không đúng",
		Code:    201,
	},
	{
		Key:     StaffStatusInactive,
		Message: "Tài khoản đăng nhập không còn hoạt động",
		Code:    202,
	},
}
