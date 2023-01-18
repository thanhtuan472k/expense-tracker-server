package route

import (
	"expense-tracker-server/external/util/routemiddleware"
	"github.com/labstack/echo/v4"
)

// Init ...
func Init(e *echo.Echo) {

	// Middlewares ...
	e.Use(routemiddleware.CORSConfig())
	e.Use(routemiddleware.Locale)

	r := e.Group("/admin/expense")

	// Components
	common(r)
	staff(r)
	staffManagement(r)
	category(r)
}
