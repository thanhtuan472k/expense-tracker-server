package route

import (
	"expense-tracker-server/pkg/app/handler"
	routevalidation "expense-tracker-server/pkg/app/route/validation"
	"github.com/labstack/echo/v4"
)

// income ...
func income(e *echo.Group) {
	var (
		g = e.Group("/incomes")
		h = handler.Income{}
		v = routevalidation.Income{}
	)

	// Create ...
	g.POST("", h.Create, v.Create)
}