package handler

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	"expense-tracker-server/pkg/admin/service"
	"github.com/labstack/echo/v4"
)

// Category ...
type Category struct{}

// Create ...
func (Category) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.CategoryBody)
		s       = service.Category()
	)

	result, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, responsemodel.ResponseCreate{ID: result}, "")
}

// All ...
func (Category) All() {
}
