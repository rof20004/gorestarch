package database

import (
	"database/sql"
	"log"
)

var database *sql.DB

func init() {
	db, err := sql.Open("sqlite3", "gorestarch.db")
	if err != nil {
		log.Fatalln(err)
	}

	// Verify connection
	if err = db.Ping(); err != nil {
		log.Fatalln(err)
	}

	database = db
}

// GetConnection - return database connection instance
func GetConnection() *sql.DB {
	return database
}

// DoMigrations - create tables
func DoMigrations() {
	var db = GetConnection()

	createTableUser := `CREATE TABLE IF NOT EXISTS users (
												id INTEGER PRIMARY KEY,
												username TEXT NOT NULL,
												password TEXT NOT NULL
											);`
	_, err := db.Exec(createTableUser)
	if err != nil {
		log.Fatalln(err)
	}
}
