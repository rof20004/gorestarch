package auth

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

const genericError = "Ocorreu um erro ao realizar a operação"
const genericSuccess = "Operação realizada com sucesso"

// Authenticate - signin user
func Authenticate(c echo.Context) error {
	auth := new(Auth)
	if err := c.Bind(auth); err != nil {
		log.Println("[Authenticate]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, genericError)
	}

	err := Get(auth)
	if err != nil {
		log.Println("[Authenticate]", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, genericError)
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	// claims["name"] = "Jon Snow"
	// claims["admin"] = false
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("gorestarch"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
