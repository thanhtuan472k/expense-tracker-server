package route

import (
	"expense-tracker-server/pkg/app/handler"
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
	g.POST("/login", h.Register)

	// GetMe ...
	g.GET("/me", h.GetMe)
}
