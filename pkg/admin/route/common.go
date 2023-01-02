package route

import (
	"expense-tracker-server/pkg/admin/handler"
	"github.com/labstack/echo/v4"
)

// common ...
func common(e *echo.Group) {
	var (
		g = e.Group("")
		h = handler.Common{}
	)

	// Ping
	g.GET("/ping", h.Ping)
}
