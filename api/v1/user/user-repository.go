package user

import (
	"database/sql"
	"errors"

	"github.com/rof20004/gorestarch/database"
)

// List - query all users
func List() ([]User, error) {
	var (
		id       sql.NullInt64
		username sql.NullString
	)

	query := "SELECT id, username FROM users"

	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		err = rows.Scan(&id, &username)
		if err != nil {
			return nil, err
		}
		users = append(users, User{ID: id.Int64, Username: username.String})
	}

	return users, nil
}

// Insert - create user
func Insert(user *User) error {
	query := "INSERT INTO users(username, password) VALUES(?, ?)"

	_, err := database.DB.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// Delete - remove user
func Delete(id string) error {
	query := "DELETE FROM users WHERE id = ?"

	result, err := database.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("Usuário não existe na base de dados => ID: " + id)
	}

	return nil
}
