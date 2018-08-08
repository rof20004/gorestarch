package user

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
)

const genericError = "Ocorreu um erro ao realizar a operação"
const genericSuccess = "Operação realizada com sucesso"

// ListUsers - list all users
func ListUsers(c echo.Context) error {
	// Exemplo de verificação de perfil de usuário com JWT
	// user := c.Get("user").(*jwt.Token)
	// claims := user.Claims.(jwt.MapClaims)
	// admin := claims["admin"].(bool)

	// if !admin {
	// 	return echo.ErrUnauthorized
	// }

	users, err := FindAll()
	if err != nil {
		log.Println("[ListUsers]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, genericError)
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser - get a user
func GetUser(c echo.Context) error {
	user, err := FindByID(c.Param("id"))
	if err != nil {
		log.Println("[GetUser]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, genericError)
	}
	return c.JSON(http.StatusOK, user)
}

// InsertUser - create a user
func InsertUser(c echo.Context) error {
	user := new(User)
	if err := c.Bind(user); err != nil {
		log.Println("[InsertUser]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, genericError)
	}

	err := Create(user)
	if err != nil {
		log.Println("[InsertUser]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, genericError)
	}

	return c.String(http.StatusOK, genericSuccess)
}

// DeleteUser - delete a user
func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	err := Delete(id)
	if err != nil {
		log.Println("[DeleteUser]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, genericError)
	}

	return c.String(http.StatusOK, genericSuccess)
}
