package handler

import (
	"expense-tracker-server/external/constant"
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	"expense-tracker-server/external/util/mgquerry"
	"expense-tracker-server/external/util/pagetoken"
	querymodel "expense-tracker-server/pkg/app/model/query"
	"expense-tracker-server/pkg/app/service"
	"github.com/labstack/echo/v4"
)

// Category ...
type Category struct{}

// All godoc
// @tags Category
// @summary All
// @id app-category-all
// @accept json
// @produce json
// @param payload body querymodel.CategoryAll true "Payload"
// @success 200 {object} nil
// @router /categories [get]
func (Category) All(c echo.Context) error {
	var (
		ctx       = echocontext.GetContext(c)
		qParams   = echocontext.GetQuery(c).(querymodel.CategoryAll)
		pageToken = pagetoken.PageTokenDecode(qParams.PageToken)
		q         = mgquerry.AppQuery{
			Page:  pageToken.Page,
			Limit: int64(constant.Limit20),
		}
		s = service.Category()
	)

	result, err := s.All(ctx, q)
	if err != nil {
		return response.R400(c, echo.Map{}, err.Error())
	}
	return response.R200(c, result, "")
}
