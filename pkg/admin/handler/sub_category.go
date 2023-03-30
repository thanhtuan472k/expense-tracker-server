package handler

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	"expense-tracker-server/pkg/admin/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SubCategory ...
type SubCategory struct{}

// Update godoc
// @tags SubCategory
// @summary Update
// @id sub-category-update
// @security ApiKeyAuth
// @accept json
// @produce json
// @Param  id path string true "Sub category id"
// @param payload body requestmodel.SubCategoryBodyCreate true "Payload"
// @success 200 {object} response.ResponseCreate
// @router /sub-categories/{id} [put]
func (SubCategory) Update(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.SubCategoryBodyUpdate)
		s       = service.SubCategory()
		id      = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)

	result, err := s.Update(ctx, id, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, response.ResponseCreate{ID: result}, "")
}

// Detail godoc
// @tags SubCategory
// @summary Detail
// @id sub-category-detail
// @security ApiKeyAuth
// @accept json
// @produce json
// @Param  id path string true "Sub category id"
// @success 200 {object} nil
// @router /sub-categories/{id} [get]
func (SubCategory) Detail(c echo.Context) error {
	var (
		ctx = echocontext.GetContext(c)
		s   = service.SubCategory()
		id  = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)

	result, err := s.Detail(ctx, id)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, result, "")
}

// ChangeStatus godoc
// @tags SubCategory
// @summary ChangeStatus
// @id sub-category-change-status
// @security ApiKeyAuth
// @accept json
// @produce json
// @Param  id path string true "Sub category id"
// @param payload body requestmodel.SubCategoryChangeStatus true "Payload"
// @success 200 {object} nil
// @router /sub-categories/{id}/status [patch]
func (SubCategory) ChangeStatus(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.SubCategoryChangeStatus)
		s       = service.SubCategory()
		id      = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)

	result, err := s.ChangeStatus(ctx, id, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, result, "")
}
