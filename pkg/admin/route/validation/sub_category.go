package validation

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	requestmodel "expense-tracker-server/pkg/admin/model/request"
	"github.com/labstack/echo/v4"
)

// SubCategory ...
type SubCategory struct{}

// Create ...
func (SubCategory) Create(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.SubCategoryBodyCreate

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}

// Update ...
func (SubCategory) Update(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var payload requestmodel.SubCategoryBodyUpdate

		if err := c.Bind(&payload); err != nil {
			return response.R400(c, nil, "")
		}

		if err := payload.Validate(); err != nil {
			return response.RouteValidation(c, err)
		}

		echocontext.SetPayload(c, payload)
		return next(c)
	}
}
