package main

import (
	"expense-tracker-server/docs/app"
	"expense-tracker-server/internal/config"
	"expense-tracker-server/pkg/app/server"
	echoSwagger "github.com/swaggo/echo-swagger"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
)

// @title Expense Tracker - Admin API
// @version 1.0
// @description All APIs for Expense app.
// @description
// @description ******************************
// @description - Add description
// @description ******************************
// @description
// @termsOfService
// @contact.name Dev team
// @contact.url
// @contact.email tuanngo.472000@gmail.com
// @basePath /app/expense

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// Echo instance
	e := echo.New()

	e.Use(apmechov4.Middleware())

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} | ${remote_ip} | ${method} ${uri} - ${status} - ${latency_human}\n",
	}))
	e.Use(middleware.Gzip())
	if os.Getenv("ENV") == "release" {
		e.Use(middleware.Recover())
	}

	// Bootstrap things
	server.Bootstrap(e)

	// Swagger
	if config.IsEnvDevelop() {
		domain := os.Getenv("DOMAIN_EXPENSE_TRACKER_APP")
		app.SwaggerInfo.Host = domain
		e.GET(app.SwaggerInfo.BasePath+"/swagger/*", echoSwagger.WrapHandler)
	}

	// Start server
	e.Logger.Fatal(e.Start(config.GetENV().Admin.Port))

}
