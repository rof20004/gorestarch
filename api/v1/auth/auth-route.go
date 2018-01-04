package auth

import (
	"github.com/labstack/echo"
)

// MapRoutePath - set routes to user
func MapRoutePath(api *echo.Group, path string) {
	group := api.Group(path)
	group.POST("", Authenticate)
}
