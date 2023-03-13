package routeauth

import (
	"expense-tracker-server/external/response"
	"expense-tracker-server/external/util/echocontext"
	"github.com/labstack/echo/v4"
)

// RequiredLogin ...
func RequiredLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := echocontext.GetCurrenUserByToken(c.Get("user"))
		if user == nil || user.ID == "" {
			return response.R403(c, echo.Map{}, response.CommonForbidden)
		}

		c.Set("user", user.ID)
		return next(c)
	}
}
