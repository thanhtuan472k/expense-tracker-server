package handler

import (
	"expense-tracker-server/external/constant"
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/external/util/pagetoken"
	"expense-tracker-server/external/util/ptime"
	responsemodel "expense-tracker-server/pkg/admin/model/response"
	querymodel "expense-tracker-server/pkg/app/model/query"
	requestmodel "expense-tracker-server/pkg/app/model/request"
	"expense-tracker-server/pkg/app/service"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IncomeMoney ...
type IncomeMoney struct{}

// Create godoc
// @tags IncomeMoney
// @summary Create
// @id app-income-money-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.IncomeMoneyBodyCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /income-moneys [post]
func (IncomeMoney) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		s       = service.IncomeMoney()
		payload = echocontext.GetPayload(c).(requestmodel.IncomeMoneyBodyCreate)
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
// @tags IncomeMoney
// @summary Update
// @id app-income-money-update
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Income money id"
// @param payload body requestmodel.IncomeMoneyBodyUpdate true "Payload"
// @success 200 {object} responsemodel.ResponseUpdate
// @router /income-moneys/{id} [put]
func (IncomeMoney) Update(c echo.Context) error {
	var (
		ctx      = echocontext.GetContext(c)
		s        = service.IncomeMoney()
		payload  = echocontext.GetPayload(c).(requestmodel.IncomeMoneyBodyUpdate)
		userID   = echocontext.GetCurrentUserID(c)
		incomeID = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)

	result, err := s.Update(ctx, userID, incomeID, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, result, "")
}

// All godoc
// @tags IncomeMoney
// @summary All
// @id app-income-money-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query querymodel.IncomeMoneyAll true "Query"
// @success 200 {object} nil
// @router /income-moneys  [get]
func (IncomeMoney) All(c echo.Context) error {
	var (
		ctx       = echocontext.GetContext(c)
		qParams   = echocontext.GetQuery(c).(querymodel.IncomeMoneyAll)
		pageToken = pagetoken.PageTokenDecode(qParams.PageToken)
		q         = mgquerry.AppQuery{
			Page:          int64(pageToken.Page),
			Limit:         int64(constant.Limit20),
			SortInterface: bson.D{{"createdAt", -1}},
			SortString:    qParams.Sort,
			ExpenseTracker: mgquerry.ExpenseTracker{
				FromAt: ptime.TimeParseISODate(qParams.FromAt),
				ToAt:   ptime.TimeParseISODate(qParams.ToAt),
			},
		}
		s      = service.IncomeMoney()
		userID = echocontext.GetCurrentUserID(c)
	)

	result, err := s.All(ctx, q, userID)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, result, "")
}
