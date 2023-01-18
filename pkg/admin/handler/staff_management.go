package handler

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	"expense-tracker-server/pkg/admin/service"
	"github.com/labstack/echo/v4"
)

// StaffManagement ...
type StaffManagement struct{}

// Create godoc
// @tags Staff
// @summary Create
// @id staff-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.StaffBodyCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /staffs [post]
func (StaffManagement) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.StaffBodyCreate)
		s       = service.Staff()
	)

	result, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, responsemodel.ResponseCreate{ID: result}, "")
}
