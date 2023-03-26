package routevalidation

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	querymodel "expense-tracker-server/pkg/app/model/query"
	"github.com/labstack/echo/v4"
)

// Category ...
type Category struct{}

// All ...
func (Category) All(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var query querymodel.CategoryAll

		if err := c.Bind(&query); err != nil {
			return response.R400(c, nil, "")
		}

		if err := query.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetQuery(c, query)
		return next(c)
	}
}
