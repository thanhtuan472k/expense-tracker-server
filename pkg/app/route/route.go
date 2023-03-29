package route

import (
	"expense-tracker-server/external/util/routemiddleware"
	routeauth "expense-tracker-server/pkg/app/route/auth"
	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo) {

	// Middlewares ...
	e.Use(routeauth.Jwt())

	e.Use(routemiddleware.CORSConfig())
	e.Use(routemiddleware.Locale)

	r := e.Group("/app/expense")

	// Components
	common(r)
	user(r)
	category(r)
	income(r)
	expenseMoney(r)
}
