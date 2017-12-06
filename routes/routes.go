package routes

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rof20004/gorestarch/api/v1/auth"
	"github.com/rof20004/gorestarch/api/v1/user"
)

// Server program
var Server = echo.New()

func init() {
	api := Server.Group("/api")

	// Auth routes
	authGroup := api.Group("/v1/auth")
	authGroup.POST("", auth.Authenticate)

	// User routes
	userGroup := api.Group("/v1/users")
	userGroup.GET("", user.ListUsers)
	userGroup.POST("", user.InsertUser, middleware.JWT([]byte("gorestarch")))
}
