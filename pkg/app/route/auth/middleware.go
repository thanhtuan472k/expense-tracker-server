package routeauth

import (
	"expense-tracker-server/external/util/echocontext"
	"expense-tracker-server/internal/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Jwt ...
func Jwt() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(config.GetENV().App.Auth.SecretKey),
		Skipper: func(c echo.Context) bool {
			token := echocontext.GetToken(c)
			user := echocontext.GetCurrenStaffByToken(c.Get("user"))
			var userID = ""
			if user != nil && user.ID != "" {
				userID = user.ID
			}

			c.Set("current_user_id", userID) // TODO: refactor bỏ vào constant
			return token == ""
		},
	})
}
