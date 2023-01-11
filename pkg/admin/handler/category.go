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

// Create godoc
// @tags Category
// @summary Create
// @id category-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.CategoryBodyCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /categories [post]
func (Category) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.CategoryBodyCreate)
		s       = service.Category()
	)

	result, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, responsemodel.ResponseCreate{ID: result}, "")
}

// Update godoc
// @tags Category
// @summary Update
// @id category-update
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Category id"
// @param payload body requestmodel.CategoryBodyCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /categories/{id} [PUT]
func (Category) Update(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.CategoryBodyCreate)
		s       = service.Category()
	)

	result, err := s.Create(ctx, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, responsemodel.ResponseCreate{ID: result}, "")
}
