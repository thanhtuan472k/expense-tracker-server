package handler

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	"expense-tracker-server/pkg/app/service"
	"github.com/labstack/echo/v4"
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
// @param payload body requestmodel.SellerSavingCampaignBodyCreate true "Payload"
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
// @summary Create
// @id create-income-money
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.SellerSavingCampaignBodyCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /incomes/{id} [put]
func (Income) Update(c echo.Context) error {
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
