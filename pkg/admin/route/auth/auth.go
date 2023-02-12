package routeauth

import (
	"fmt"
	"github.com/labstack/echo/v4"
)

// RequiredLogin ...
func RequiredLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		staffToken := c.Request().Header.Get("token")

		fmt.Println("staffToken", staffToken)
		return next(c)
	}
}
