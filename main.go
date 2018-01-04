package main

import (
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rof20004/gorestarch/database"
	"github.com/rof20004/gorestarch/routes"
)

func main() {
	e := echo.New()

	// API initialize
	api := e.Group("/api")

	// ROUTES initialize
	routes.Initialize(api)

	// DATABASE initialize
	database.DoMigrations()

	// PROGRAM initialize
	e.Logger.Fatal(e.Start(":8080"))
}
