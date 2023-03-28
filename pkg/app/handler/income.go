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

// Income ...
type Income struct{}

// Create godoc
// @tags Income
// @summary Create
// @id app-income-money-create
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
// @id app-income-money-update
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Income money id"
// @param payload body requestmodel.IncomeBodyUpdate true "Payload"
// @success 200 {object} responsemodel.ResponseUpdate
// @router /incomes/{id} [put]
func (Income) Update(c echo.Context) error {
	var (
		ctx      = echocontext.GetContext(c)
		s        = service.Income()
		payload  = echocontext.GetPayload(c).(requestmodel.IncomeBodyUpdate)
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
// @tags Income
// @summary All
// @id app-income-money-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query querymodel.IncomeMoneyAll true "Query"
// @success 200 {object} nil
// @router /incomes [get]
func (Income) All(c echo.Context) error {
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
		s      = service.Income()
		userID = echocontext.GetCurrentUserID(c)
	)

	result, err := s.All(ctx, q, userID)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, result, "")
}
