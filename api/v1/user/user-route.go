package user

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// MapRoutePath - set routes to user
func MapRoutePath(e *echo.Group, path string) {
	group := e.Group(path)
	group.GET("", ListUsers)
	group.POST("", InsertUser, middleware.JWT([]byte("gorestarch")))
}
