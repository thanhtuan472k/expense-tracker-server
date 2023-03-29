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

// ExpenseMoney ...
type ExpenseMoney struct{}

// Create godoc
// @tags ExpenseMoney
// @summary Create
// @id app-expense-money-create
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload body requestmodel.ExpenseMoneyBodyCreate true "Payload"
// @success 200 {object} responsemodel.ResponseCreate
// @router /expense-moneys [post]
func (ExpenseMoney) Create(c echo.Context) error {
	var (
		ctx     = echocontext.GetContext(c)
		s       = service.ExpenseMoney()
		payload = echocontext.GetPayload(c).(requestmodel.ExpenseMoneyBodyCreate)
		userID  = echocontext.GetCurrentUserID(c)
	)

	expenseID, err := s.Create(ctx, userID, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, responsemodel.ResponseCreate{
		ID: expenseID,
	}, "")
}

// Update godoc
// @tags ExpenseMoney
// @summary Update
// @id app-expense-money-update
// @security ApiKeyAuth
// @accept json
// @produce json
// @param id path string true "Expense money id"
// @param payload body requestmodel.ExpenseMoneyBodyUpdate true "Payload"
// @success 200 {object} responsemodel.ResponseUpdate
// @router /expense-moneys/{id} [put]
func (ExpenseMoney) Update(c echo.Context) error {
	var (
		ctx       = echocontext.GetContext(c)
		s         = service.ExpenseMoney()
		payload   = echocontext.GetPayload(c).(requestmodel.ExpenseMoneyBodyUpdate)
		userID    = echocontext.GetCurrentUserID(c)
		expenseID = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)

	result, err := s.Update(ctx, userID, expenseID, payload)
	if err != nil {
		return response.R400(c, nil, err.Error())
	}

	return response.R200(c, result, "")
}

// All godoc
// @tags ExpenseMoney
// @summary All
// @id app-expense-money-all
// @security ApiKeyAuth
// @accept json
// @produce json
// @param payload query querymodel.ExpenseMoneyAll true "Query"
// @success 200 {object} nil
// @router /expense-moneys  [get]
func (ExpenseMoney) All(c echo.Context) error {
	var (
		ctx       = echocontext.GetContext(c)
		qParams   = echocontext.GetQuery(c).(querymodel.ExpenseMoneyAll)
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
		s      = service.ExpenseMoney()
		userID = echocontext.GetCurrentUserID(c)
	)

	result, err := s.All(ctx, q, userID)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, result, "")
}

// Detail godoc
// @tags ExpenseMoney
// @summary Detail
// @id app-expense-money-detail
// @security ApiKeyAuth
// @accept json
// @produce json
// @Param  id path string true "Expense money id"
// @success 200 {object} nil
// @router /expense-moneys/{id} [get]
func (ExpenseMoney) Detail(c echo.Context) error {
	var (
		ctx = echocontext.GetContext(c)
		s   = service.ExpenseMoney()
		id  = echocontext.GetParam(c, "id").(primitive.ObjectID)
	)

	result, err := s.Detail(ctx, id)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, result, "")
}
