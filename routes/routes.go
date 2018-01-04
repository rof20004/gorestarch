package routes

import (
	"github.com/labstack/echo"
	"github.com/rof20004/gorestarch/api/v1/auth"
	"github.com/rof20004/gorestarch/api/v1/user"
)

// Initialize and configure application routes
func Initialize(api *echo.Group) {
	auth.MapRoutePath(api, "/v1/auth")
	user.MapRoutePath(api, "/v1/users")
}
