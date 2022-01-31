package middlewares

import (
	"log"
	"net/http"
	"os"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/handler"
	"github.com/achange8/learnecho2/models"
	"github.com/dgrijalva/jwt-go"
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
		if err == nil {
			///auth access token
			rawtoken := cookie.Value
			_, err = jwt.ParseWithClaims(rawtoken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("key")), nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					return c.JSON(http.StatusUnauthorized, "signature is invalid")
				}
				return c.JSON(http.StatusForbidden, "Authorization failed")
			}
			return next(c)
		} else {
			//recreate access token if exist reftokn.
			db := db.Connect()
			refresh := new(models.Refresh)
			cookie, err := c.Cookie("JWTRefreshToken")
			if err != nil { //todo go signin point
				return c.JSON(http.StatusUnauthorized, "no refresh token! signin again")
			}
			///// auth & parse refresh token ///
			rawRefreshtoken := cookie.Value
			refreshClaims := &jwt.StandardClaims{}
			_, err = jwt.ParseWithClaims(
				rawRefreshtoken, refreshClaims,
				func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("key2")), nil
				})
			if err != nil {
				return c.JSON(http.StatusUnauthorized, "login again")
			}
			username := refreshClaims.Id //<-parsed username in jwt
			//dont reftoken in db
			result := db.Find(&refresh, "reftoken=?", rawRefreshtoken)
			if result.RowsAffected == 0 { //todo go signin point
				return c.JSON(http.StatusUnauthorized, "Do signin again")
			}
			newtoken, err := handler.CreateAccessToken(username)
			if err != nil {
				return c.JSON(401, "failed create new token")
			}
			JWTaccessCookie := handler.CreateAccessCookie(username, newtoken)
			c.SetCookie(JWTaccessCookie)

			_, err = jwt.ParseWithClaims(newtoken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
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
}
