package auth

import (
	"database/sql"

	"github.com/rof20004/gorestarch/database"
)

// Get - get user credentials
func Get(auth *Auth) error {
	query := "SELECT id FROM users WHERE username = ? and password = ?"

	var id sql.NullInt64

	err := database.DB.QueryRow(query, auth.Username, auth.Password).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}
