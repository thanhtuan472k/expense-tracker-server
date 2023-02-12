package route

import (
	"expense-tracker-server/pkg/admin/handler"
	routeauth "expense-tracker-server/pkg/admin/route/auth"
	"expense-tracker-server/pkg/admin/route/routevalidation"
	"github.com/labstack/echo/v4"
)

// staff ...
func staff(e *echo.Group) {
	var (
		g = e.Group("/staffs")
		h = handler.Staff{}
		v = routevalidation.Staff{}
	)

	// Login ...
	g.POST("/login", h.Login, v.Login)

	// GetMe ...
	g.GET("/me", h.GetMe, routeauth.RequiredLogin)

	// Update ...
	g.PUT("/me", h.Update, v.Update)

	// ChangePassword ...
	g.PATCH("/me/password", h.ChangePassword)
}
