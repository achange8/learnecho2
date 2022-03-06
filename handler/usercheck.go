package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

//return logged in user if not logged in, return nil
func Usercheck(c echo.Context) error {
	envERR := godotenv.Load(".env")
	if envERR != nil {
		log.Println("Could not load .env file")
		os.Exit(1)
	}
	cookie, err := c.Cookie("JWTaccessToken")
	if err == nil {
		///auth access token
		rawtoken := cookie.Value
		claims := &jwt.StandardClaims{}
		_, err = jwt.ParseWithClaims(rawtoken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("key")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.JSON(http.StatusUnauthorized, "signature is invalid")
			}
			return c.JSON(http.StatusForbidden, "Authorization failed")
		}
		username := claims.Id
		return c.JSON(http.StatusOK, username)
	}
	return nil
}
