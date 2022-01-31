package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func Checktoken(c echo.Context) error {
	envERR := godotenv.Load(".env")
	if envERR != nil {
		log.Println("Could not load .env file")
		os.Exit(1)
	}
	cookie, err := c.Cookie("JWTRefreshToken")
	if err != nil {
		return err
	}
	rawRefreshtoken := cookie.Value

	///// auth & parse refresh token ///
	claims := &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(rawRefreshtoken, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("key2")), nil
		})
	if err != nil {
		return c.JSON(http.StatusUnauthorized, "login again")
	} else {
		username := claims.Id
		return c.JSON(http.StatusOK, username)
	}
}
