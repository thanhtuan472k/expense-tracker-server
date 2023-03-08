package server

import (
	"expense-tracker-server/external/logger"
	"expense-tracker-server/pkg/app/route"
	"expense-tracker-server/pkg/app/server/initialize"
	"github.com/labstack/echo/v4"
)

// Bootstrap ...
func Bootstrap(e *echo.Echo) {
	logger.Init("expense", "expense-app")

	// Init modules
	initialize.Init()

	// Routes
	route.Init(e)
}
