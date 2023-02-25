package route

import (
	"expense-tracker-server/pkg/admin/handler"
	routeauth "expense-tracker-server/pkg/admin/route/auth"
	"expense-tracker-server/pkg/admin/route/validation"
	"github.com/labstack/echo/v4"
)

// subCategory ...
func subCategory(e *echo.Group) {
	var (
		g = e.Group("/sub-categories", routeauth.RequiredLogin)
		h = handler.SubCategory{}
		v = validation.SubCategory{}
	)

	// Update ...
	g.PUT("/:id", h.Update, v.Update, v.Detail)

	// Detail ...
	g.GET("/:id", h.Detail, v.Detail)

	// ChangeStatus ...
	g.PATCH("/:id/status", h.ChangeStatus, v.ChangeStatus, v.Detail)
}
