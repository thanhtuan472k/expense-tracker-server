package route

import (
	"expense-tracker-server/pkg/admin/handler"
	"expense-tracker-server/pkg/admin/route/routevalidation"
	"github.com/labstack/echo/v4"
)

// category ...
func category(e *echo.Group) {
	var (
		g = e.Group("/categories")

		h = handler.Category{}
		v = routevalidation.Category{}
	)

	// Create ...
	g.POST("", h.Create, v.ValidateBody)

	// All ...
	g.GET("", h.All, v.All)

	// Detail ...
	g.GET("/:id", h.Detail, v.Detail)

	// Update ...
	g.PUT("/:id", h.Update, v.ValidateBody)

	// ChangeStatus ...
	g.PATCH("/:id", h.ChangeStatus, v.ValidateBody, v.Detail)
}