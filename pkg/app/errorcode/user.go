package errorcode

import "expense-tracker-server/external/response"

const (
	UserExistedPhoneOrEmail        = "user_existed_phone_or_email"
	UserNotFound                   = "user_not_found"
	UserPasswordIncorrect          = "user_password_incorrect" // login
	UserStatusInactive             = "user_status_inactive"
	UserWrongOldPassword           = "user_wrong_old_password"             // khi nhập sai mật khẩu cũ
	UserNewPasswordSameOldPassword = "user_new_password_same_old_password" // khi mật khẩu mới giống mật khẩu cũ
)

// 200 - 299
var user = []response.Code{
	{
		Key:     UserExistedPhoneOrEmail,
		Message: "Số điện thoại hoặc email đã tồn tại",
		Code:    200,
	},
	{
		Key:     UserNotFound,
		Message: "Không tìm thấy tài khoản người dùng ",
		Code:    201,
	},
	{
		Key:     UserPasswordIncorrect,
		Message: "Mật khẩu không đúng",
		Code:    202,
	},
	{
		Key:     UserStatusInactive,
		Message: "Tài khoản đăng nhập không còn hoạt động",
		Code:    203,
	},
	{
		Key:     UserWrongOldPassword,
		Message: "Bạn đã nhập sai mật khẩu cũ. Vui lòng nhập lại! ",
		Code:    203,
	},
	{
		Key:     UserNewPasswordSameOldPassword,
		Message: "Mật khẩu mới giống với mật khẩu cũ. Vui lòng nhập lại! ",
		Code:    204,
	},
}
