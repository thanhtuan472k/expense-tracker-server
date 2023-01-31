package handler

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	"expense-tracker-server/pkg/admin/service"
	"github.com/labstack/echo/v4"
)

// Staff ...
type Staff struct{}

// Login godoc
// @tags Staff
// @summary Login
// @id staff-login
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.StaffBodyLogin true "Payload"
// @success 200 {object} responsemodel.ResponseLoginSuccess
// @router /staffs [post]
func (Staff) Login(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.StaffBodyLogin)
		s       = service.Staff()
	)

	result, err := s.Login(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, result, "")
}
