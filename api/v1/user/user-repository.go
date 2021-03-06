package user

import (
	"database/sql"
	"errors"
	"strconv"

	"github.com/rof20004/gorestarch/database"
)

// FindAll - query all users
func FindAll() ([]User, error) {
	var db = database.GetConnection()

	var (
		id       sql.NullInt64
		username sql.NullString
	)

	query := "SELECT id, username FROM users"

	rows, err := db.Query(query)
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

// FindByID - query a user
func FindByID(id string) (*User, error) {
	var db = database.GetConnection()

	var (
		username sql.NullString
	)

	query := "SELECT username FROM users WHERE id = ?"

	err := db.QueryRow(query, id).Scan(&username)
	if err != nil {
		return nil, err
	}

	i, _ := strconv.ParseInt(id, 0, 64)

	return &User{
		ID:       i,
		Username: username.String,
	}, nil
}

// Create - create user
func Create(user *User) error {
	var db = database.GetConnection()

	query := "INSERT INTO users(username, password) VALUES(?, ?)"

	_, err := db.Exec(query, user.Username, user.Password)
	if err != nil {
		return err
	}

	return nil
}

// Delete - remove user
func Delete(id string) error {
	var db = database.GetConnection()

	query := "DELETE FROM users WHERE id = ?"

	result, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		return errors.New("Usuário não existe na base de dados => ID: " + id)
	}

	return nil
}
