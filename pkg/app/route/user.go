package route

import (
	"expense-tracker-server/pkg/app/handler"
	routeauth "expense-tracker-server/pkg/app/route/auth"
	routevalidation "expense-tracker-server/pkg/app/route/validation"
	"github.com/labstack/echo/v4"
)

// user ...
func user(e *echo.Group) {
	var (
		g = e.Group("/users")
		h = handler.User{}
		v = routevalidation.User{}
	)

	// Register ...
	g.POST("/register", h.Register, v.Register)

	// Login ...
	g.POST("/login", h.Login, v.Login)

	// GetMe ...
	g.GET("/me", h.GetMe, routeauth.RequiredLogin)
}
