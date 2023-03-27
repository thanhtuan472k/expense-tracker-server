package route

import (
	"expense-tracker-server/pkg/app/handler"
	routeauth "expense-tracker-server/pkg/app/route/auth"
	routevalidation "expense-tracker-server/pkg/app/route/validation"
	"github.com/labstack/echo/v4"
)

// income ...
func income(e *echo.Group) {
	var (
		g = e.Group("/incomes", routeauth.RequiredLogin)
		h = handler.Income{}
		v = routevalidation.Income{}
	)

	// Create ...
	g.POST("", h.Create, v.Create)

	// Update ...
	g.PUT("/:id", h.Update, v.Update)
}
