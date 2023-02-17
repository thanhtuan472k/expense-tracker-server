package routeauth

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	"github.com/labstack/echo/v4"
)

// RequiredLogin ...
func RequiredLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check invalid token
		staff := echocontext.GetCurrenStaffByToken(c.Get("user"))

		if staff == nil || staff.ID == "" {
			return response.R403(c, echo.Map{}, response.CommonForbidden)
		}

		staffResponse := echocontext.Staff{
			ID:   staff.ID,
			Name: staff.Name,
		}

		c.Set("staff", staffResponse)
		return next(c)
	}
}
