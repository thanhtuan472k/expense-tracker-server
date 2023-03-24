package route

import (
	"expense-tracker-server/pkg/app/handler"
	routevalidation "expense-tracker-server/pkg/app/route/validation"
	"github.com/labstack/echo/v4"
)

// category ...
func category(e *echo.Group) {
	var (
		g = e.Group("/categories")
		h = handler.Category{}
		v = routevalidation.Category{}
	)

	// All ...
	g.GET("", h.All, v.All)
}
