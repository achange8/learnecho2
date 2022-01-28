package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/helper"
	"github.com/achange8/learnecho2/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func SignIn(c echo.Context) error {
	user := new(models.User)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	inputpw := user.Password

	db := db.Connect()
	result := db.Find(user, "email=?", user.Email)

	// 존재하지않는 아이디일 경우
	if result.RowsAffected == 0 {
		return echo.ErrBadRequest
	}

	res := helper.CheckPasswordHash(user.Password, inputpw)

	// 비밀번호 검증에 실패한 경우
	if !res {
		return echo.ErrUnauthorized
	} else {
		userid := user.Email                          //save user email -> user id
		AccessToken, err := CreateAccessToken(userid) //create ac token
		if err != nil {
			log.Println("Err Creating Access_Token!", err)
		}
		JWTaccessCookie := new(http.Cookie)
		JWTaccessCookie.Name = "JWTaccessCookie"
		JWTaccessCookie.Value = AccessToken
		JWTaccessCookie.Expires = time.Now().Add(15 * time.Minute)
		JWTaccessCookie.HttpOnly = true

		c.SetCookie(JWTaccessCookie)

		RefreshToken, err := createRefreshToken(userid) //create rf token
		if err != nil {
			log.Println("Err Creating Access_Token!", err)
		}
		RefreshCookie := new(http.Cookie)
		RefreshCookie.Name = "JWTaccessCookie"
		RefreshCookie.Value = RefreshToken
		RefreshCookie.Expires = time.Now().Add(24 * 7 * time.Hour)
		RefreshCookie.HttpOnly = true

		c.SetCookie(JWTaccessCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message":       "You were logged in!",
			"Access_Token":  AccessToken,
			"Refresh_Token": RefreshToken,
		})
	}
}

func CreateAccessToken(userID string) (string, error) {
	claims := models.JwtClaims{
		jwt.StandardClaims{
			Id:        userID,
			ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(os.Getenv("key")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func createRefreshToken(userID string) (string, error) {
	claims := models.JwtClaims{
		jwt.StandardClaims{
			Id:        userID,
			ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(os.Getenv("key")))
	if err != nil {
		return "", err
	}
	return token, nil
}
