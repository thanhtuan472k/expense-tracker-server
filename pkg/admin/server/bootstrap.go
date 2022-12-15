package server

import (
	"expense-tracker-server/external/logger"
	"expense-tracker-server/pkg/admin/route"
	"expense-tracker-server/pkg/admin/server/initialize"
	"github.com/labstack/echo/v4"
)

// Bootstrap ...
func Bootstrap(e *echo.Echo) {
	logger.Init("expense", "expense-admin")

	// Init modules
	initialize.Init()

	// Routes
	route.Init(e)
}
