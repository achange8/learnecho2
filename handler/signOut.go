package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

//check db refresh token, parse id
// recreate cookie ac,rf -time
//delete refresh token in db date
func SignOut(c echo.Context) error {
	db := db.Connect()
	refresh := new(models.Refresh)
	envERR := godotenv.Load(".env")
	if envERR != nil {
		log.Println("Could not load .env file")
		os.Exit(1)
	}

	cookie, err := c.Cookie("JWTRefreshToken")
	if err != nil {
		return err
	}
	token := cookie.Value
	refreshClaims := &jwt.StandardClaims{}
	_, err = jwt.ParseWithClaims(token, refreshClaims,
		func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("key2")), nil
		})
	if err != nil {
		c.JSON(http.StatusUnauthorized, "u didnt login")
	}
	result := db.Find(&refresh, "reftoken=?", token)
	if result.RowsAffected != 0 { //todo go signin point
		db.Where("reftoken = ?", token).Delete(&refresh)
	}
	userid := refreshClaims.Id
	JWTaccessToken := createAccesslogoutCookie(userid)
	JWTRefreshToken := createRefreshlogoutCookie(userid)
	c.SetCookie(JWTaccessToken)
	c.SetCookie(JWTRefreshToken)
	return c.JSON(http.StatusOK, "good bye")
}

func createAccesslogoutCookie(userID string) *http.Cookie {
	JWTaccessCookie := new(http.Cookie)
	JWTaccessCookie.Name = "JWTaccessToken"
	JWTaccessCookie.Value = ""
	JWTaccessCookie.Expires = time.Now().Add(-time.Hour)
	JWTaccessCookie.HttpOnly = true
	JWTaccessCookie.Path = "/"
	return JWTaccessCookie
}
func createRefreshlogoutCookie(userID string) *http.Cookie {
	RefreshCookie := new(http.Cookie)
	RefreshCookie.Name = "JWTRefreshToken"
	RefreshCookie.Value = ""
	RefreshCookie.Expires = time.Now().Add(-time.Hour)
	RefreshCookie.Path = "/"
	RefreshCookie.HttpOnly = true
	return RefreshCookie
}
