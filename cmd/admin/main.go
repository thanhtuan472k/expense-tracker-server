package main

import (
	"expense-tracker-server/internal/config"
	"expense-tracker-server/pkg/admin/server"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.elastic.co/apm/module/apmechov4"
)

// @title Selly Campaign - Admin API
// @version 1.0
// @description All APIs for Campaign admin.
// @description
// @description ******************************
// @description - Add description
// @description ******************************
// @description
// @termsOfService https://selly.vn
// @contact.name Dev team
// @contact.url https://selly.vn
// @contact.email dev@selly.vn
// @basePath /admin/campaign

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
	//if config.IsEnvDevelop() {
	//	domain := os.Getenv("DOMAIN_CAMPAIGN_ADMIN")
	//	admin.SwaggerInfo.Host = domain
	//	e.GET(admin.SwaggerInfo.BasePath+"/swagger/*", echoSwagger.WrapHandler)
	//}

	// Start server
	e.Logger.Fatal(e.Start(config.GetENV().Admin.Port))
}
