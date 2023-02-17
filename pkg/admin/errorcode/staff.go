package errorcode

import "expense-tracker-server/external/response"

const (
	StaffNotFound                   = "staff_not_found"
	StaffPasswordIncorrect          = "staff_password_incorrect" // login
	StaffStatusInactive             = "staff_status_inactive"
	StaffWrongOldPassword           = "staff_wrong_old_password"             // khi nhập sai mật khẩu cũ
	StaffNewPasswordSameOldPassword = "staff_new_password_same_old_password" // khi mật khẩu mới giống mật khẩu cũ
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
	{
		Key:     StaffWrongOldPassword,
		Message: "Bạn đã nhập sai mật khẩu cũ. Vui lòng nhập lại! ",
		Code:    203,
	},
	{
		Key:     StaffNewPasswordSameOldPassword,
		Message: "Mật khẩu mới giống với mật khẩu cũ. Vui lòng nhập lại! ",
		Code:    204,
	},
}
