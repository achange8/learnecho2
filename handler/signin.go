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
	db := db.Connect()
	user := new(models.User)
	refresh := new(models.Refresh)

	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	inputpw := user.Password
	result := db.Find(user, "id=?", user.Id)
	// 존재하지않는 아이디일 경우
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusBadRequest, "ID doesn't exist")
	}
	res := helper.CheckPasswordHash(user.Password, inputpw)

	// 비밀번호 검증에 실패한 경우
	if !res {
		return c.JSON(http.StatusUnauthorized, "wrong password")
	} else {
		//create ac token
		AccessToken, err := CreateAccessToken(user.Id)
		if err != nil {
			log.Println("Err Creating Access Token!", err)
		}
		//create ac cookie
		JWTaccessCookie := CreateAccessCookie(user.Id, AccessToken)
		c.SetCookie(JWTaccessCookie)

		////TODO : save id and Refresh token in DB///
		RefreshToken, err := createRefreshToken(user.Id)
		if err != nil {
			log.Println("Err Creating Refresh Token!", err)
		}
		id := db.Find(&refresh, "id=?", user.Id)
		if id.RowsAffected != 0 {
			db.Model(&refresh).Where("id =?", user.Id).Update("reftoken", RefreshToken)
			// UPDATE refreshes SET `reftoken` = RefreshToken WHERE id = user.Id
		} else {
			refresh.Id = user.Id
			refresh.Reftoken = RefreshToken
			db.Create(&refresh)
		}
		//make  refresh cookie
		RefreshCookie := createRefreshCookie(user.Id, RefreshToken)
		c.SetCookie(RefreshCookie)

		return c.JSON(http.StatusOK, map[string]string{
			"message":       "You were logged in!",
			"Access_Token":  AccessToken,
			"Refresh_Token": RefreshToken,
		})
	}
}

func CreateAccessCookie(userID string, AccessToken string) *http.Cookie {
	JWTaccessCookie := new(http.Cookie)
	JWTaccessCookie.Name = "JWTaccessToken"
	JWTaccessCookie.Value = AccessToken
	JWTaccessCookie.Expires = time.Now().Add(15 * time.Minute)
	JWTaccessCookie.HttpOnly = true
	JWTaccessCookie.Path = "/"
	return JWTaccessCookie
}
func createRefreshCookie(userID string, RefreshToken string) *http.Cookie {
	RefreshCookie := new(http.Cookie)
	RefreshCookie.Name = "JWTRefreshToken"
	RefreshCookie.Value = RefreshToken
	RefreshCookie.Expires = time.Now().Add(24 * 7 * time.Hour)
	RefreshCookie.Path = "/"
	RefreshCookie.HttpOnly = true
	return RefreshCookie
}

func CreateAccessToken(userID string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        userID,
		ExpiresAt: time.Now().Add(15 * time.Minute).Unix(),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(os.Getenv("key")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func createRefreshToken(userID string) (string, error) {
	claims := jwt.StandardClaims{
		Id:        userID,
		ExpiresAt: time.Now().Add(24 * 7 * time.Hour).Unix(),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err := rawToken.SignedString([]byte(os.Getenv("key2")))
	if err != nil {
		return "", err
	}
	return token, nil
}
