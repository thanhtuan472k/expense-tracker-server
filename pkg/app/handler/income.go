package handler

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	"expense-tracker-server/pkg/app/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Income ...
type Income struct{}

// Create godoc
// @tags Income
// @summary Create
// @id create-income-money
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.IncomeBodyCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /incomes [post]
func (Income) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		s       = service.Income()
		payload = echocontext.GetPayload(c).(requestmodel.IncomeBodyCreate)
		userID  = echocontext.GetCurrentUserID(c)
	)

	incomeID, err := s.Create(ctx, userID, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, responsemodel.ResponseCreate{
		ID: incomeID,
	}, "")
}

// Update godoc
// @tags Income
// @summary Update
// @id update-income-money
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Income money id"
// @param payload body requestmodel.IncomeBodyUpdate true "Payload"
// @success 200 {object} nil
// @router /incomes/{id} [put]
func (Income) Update(c echo.Context) error {
	var (
		ctx      = echocontext.GetContext(c)
		s        = service.Income()
		payload  = echocontext.GetPayload(c).(requestmodel.IncomeBodyUpdate)
		userID   = echocontext.GetCurrentUserID(c)
		incomeID = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)
	fmt.Println("userID", userID)
	//fmt.Println("incomID", incomeID)
	result, err := s.Update(ctx, userID, incomeID, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, result, "")
}
