package handler

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	"expense-tracker-server/external/util/mgquerry"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	"expense-tracker-server/pkg/admin/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

// All godoc
// @tags Category
// @summary All
// @id category-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query requestmodel.CategoryAll true "Query"
// @success 200 {object} responsemodel.ResponseCategoryAll
// @router /categories [get]
func (Category) All(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		qParams = echocontext.GetQuery(c).(requestmodel.CategoryAll)
		s       = service.Category()
		q       = mgquerry.AppQuery{
			Page:  qParams.Page,
			Limit: qParams.Limit,
			SortInterface: bson.D{
				{"createdAt", -1},
			},
		}
	)

	result := s.All(ctx, q)
	return response.R200(c, result, "")
}

// Detail godoc
// @tags Category
// @summary Detail
// @id category-detail
// @security ApiKeyAuth
// @accept json
// @produce json
// @Param  id path string true "Category id"
// @success 200 {object} responsemodel.ResponseCategoryAdmin
// @router /categories/{id} [get]
func (Category) Detail(c echo.Context) error {
	var (
		ctx = echocontext.GetContext(c)
		s   = service.Category()
		id  = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)

	result, err := s.Detail(ctx, id)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, result, "")
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
// @success 200 {object} responsemodel.ResponseUpdate
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

// ChangeStatus godoc
// @tags Category
// @summary ChangeStatus
// @id category-change-status
// @security ApiKeyAuth
// @accept json
// @produce json
// @Param  id path string true "Category id"
// @param payload body requestmodel.CategoryChangeStatus true "Payload"
// @success 200 {object} responsemodel.ResponseChangeStatus
// @router /categories/{id}/status [patch]
func (Category) ChangeStatus(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		payload = echocontext.GetPayload(c).(requestmodel.CategoryChangeStatus)
		s       = service.Category()
		id      = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)

	result, err := s.ChangeStatus(ctx, id, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}
	return response.R200(c, result, "")
}
