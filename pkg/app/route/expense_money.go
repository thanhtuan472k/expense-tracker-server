package route

import (
	"expense-tracker-server/pkg/app/handler"
	routeauth "expense-tracker-server/pkg/app/route/auth"
	routevalidation "expense-tracker-server/pkg/app/route/validation"
	"github.com/labstack/echo/v4"
)

// expenseMoney ...
func expenseMoney(e *echo.Group) {
	var (
		g = e.Group("/expense-moneys", routeauth.RequiredLogin)
		h = handler.ExpenseMoney{}
		v = routevalidation.ExpenseMoney{}
	)

	// Create ...
	g.POST("", h.Create, v.Create)

	// Update ...
	g.PUT("/:id", h.Update, v.Detail)

	// All ...
	g.GET("", h.All, v.All)

	// Detail ...
	g.GET("/:id", h.Detail, v.Detail)
}
