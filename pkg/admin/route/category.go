package route

import (
	"expense-tracker-server/pkg/admin/handler"
	routeauth "expense-tracker-server/pkg/admin/route/auth"
	"expense-tracker-server/pkg/admin/route/validation"
	"github.com/labstack/echo/v4"
)

// category ...
func category(e *echo.Group) {
	var (
		g = e.Group("/categories", routeauth.RequiredLogin)

		h  = handler.Category{}
		v  = validation.Category{}
		vs = validation.SubCategory{}
	)

	// Create ...
	g.POST("", h.Create, v.Create)

	// All ...
	g.GET("", h.All, v.All)

	// Detail ...
	g.GET("/:id", h.Detail, v.Detail)

	// Update ...
	g.PUT("/:id", h.Update, v.Update, v.Detail)

	// ChangeStatus ...
	g.PATCH("/:id/status", h.ChangeStatus, v.ChangeStatus, v.Detail)

	// CreateSubCategory ...
	g.POST("/:id/sub-categories", h.CreateSubCategory, vs.Create, v.Detail)
}
