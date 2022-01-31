package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/handler"
	"github.com/achange8/learnecho2/models"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func TokenchekMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		envERR := godotenv.Load(".env")
		if envERR != nil {
			log.Println("Could not load .env file")
			os.Exit(1)
		}
		cookie, err := c.Cookie("JWTaccessToken")
		rawtoken := cookie.Value
		//recreate access token if exist reftokn.
		if rawtoken == "" || err != nil {
			db := db.Connect()
			refresh := new(models.Refresh)
			RFcookie, err := c.Cookie("RefreshCookie")
			if err != nil { //todo go signin point
				return err
			}
			refreshtoken := RFcookie.Value
			if refreshtoken == "" { //todo go signin point
				return c.JSON(http.StatusUnauthorized, "null RF cookie Do signin again")
			}
			//dont reftoken in db
			result := db.Find(refresh, "reftoken = ?", refreshtoken)
			if result.RowsAffected == 0 { //todo go signin point
				return c.JSON(http.StatusUnauthorized, "Do signin again")
			}

			//	refresh token OK ,remake//
			token, err := jwt.Parse(refreshtoken, nil)
			if token == nil {
				return err
			}
			claims, _ := token.Claims.(jwt.MapClaims)
			userid := claims["jti"].(string)
			newtoken, err := handler.CreateAccessToken(userid)
			if err != nil {
				log.Println("Err Creating Refresh Token!", err)
			}
			JWTaccessCookie := handler.CreateAccessCookie(userid, newtoken)
			c.SetCookie(JWTaccessCookie)
		}

		///auth token
		_, err = jwt.ParseWithClaims(rawtoken, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("key")), nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				return c.JSON(http.StatusUnauthorized, "signature is invalid")
			}
			return c.JSON(http.StatusForbidden, "Authorization failed")
		}
		return next(c)
	}
}

/*
func retakecookies(c echo.Context) error {
	db := db.Connect()
	refresh := new(models.Refresh)
	cookie, err := c.Cookie("refreshtoken")
	if err != nil {
		return c.JSON(http.StatusBadRequest, "You dont have rfcookie Do signin again")
	}
	rawtoken := cookie.Value
	if rawtoken == "" {
		return c.JSON(http.StatusUnauthorized, "null RF cookie Do signin again")
	}
	result := db.Select("id").Find(&refresh, rawtoken)
	//dont reftoken in db
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusUnauthorized, "Do signin again")
	}
	token, err := jwt.Parse(rawtoken, nil)
	if token == nil {
		return err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	userid := claims["jti"].(string)
	newtoken, err := handler.CreateAccessToken(userid)
	if err != nil {
		log.Println("Err Creating Refresh Token!", err)
	}
	JWTaccessCookie := handler.CreateAccessCookie(userid, newtoken)
	c.SetCookie(JWTaccessCookie)
	return nil
}*/
