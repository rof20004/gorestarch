package user

import (
	"github.com/labstack/echo"
)

// MapRoutePath - set routes to user
func MapRoutePath(e *echo.Group, path string) {
	group := e.Group(path)
	group.GET("", ListUsers)
	group.GET("/:id", GetUser)
	group.POST("", InsertUser)
}
