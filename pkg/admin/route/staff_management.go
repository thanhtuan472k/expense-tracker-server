package route

import (
	"expense-tracker-server/pkg/admin/handler"
	"expense-tracker-server/pkg/admin/route/routevalidation"
	"github.com/labstack/echo/v4"
)

// staffManagement ...
func staffManagement(e *echo.Group) {
	var (
		g = e.Group("/staffs")
		h = handler.StaffManagement{}
		v = routevalidation.StaffManagement{}
	)

	// Create ... --> Just admin can create staff
	g.POST("", h.Create, v.Create)
}
