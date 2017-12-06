package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/rof20004/gorestarch/database"
	"github.com/rof20004/gorestarch/routes"
)

func main() {
	database.InitDatabase()
	database.InitMigrations()
	routes.Server.Logger.Fatal(routes.Server.Start(":8080"))
}
