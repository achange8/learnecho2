package router2

import (
	"log"
	"net/http"
	"os"

	"github.com/achange8/learnecho2/handler"
	"github.com/achange8/learnecho2/middlewares"
	"github.com/golang-jwt/jwt"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	g := e.Group("/user")
	w := e.Group("/board")
	w.Use(middlewares.TokenchekMiddleware)
	middlewares.Setlogincheker(g)
	e.GET("/ckecktoken", checktoken)
	e.GET("/api/signin", handler.SignIn)
	w.GET("/write", handler.Boardform)
	g.GET("/profile", Uprofile)
	w.POST("/write", handler.WriteBoard) //upload board
	e.POST("/api/signup", handler.Signup)
	return e
}

func checktoken(c echo.Context) error {
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

func Uprofile(c echo.Context) error {
	// user := c.Get("user")
	// token := user.(*jwt.Token)
	// claims := token.Claims.(jwt.MapClaims)
	// log.Println("User ID: ", claims["jti"])

	return c.String(http.StatusOK, "logged in user page")
}

//accesstoken checker
// func tokencheker(c echo.Context) error {
// 	envERR := godotenv.Load(".env")
// 	if envERR != nil {
// 		log.Println("Could not load .env file")
// 		os.Exit(1)
// 	}
// 	cookie, err := c.Cookie("JWTaccessToken")
// 	if err != nil { //TODO : reqeust refresh token ,create new actoken or login again
// 		return c.JSON(http.StatusBadRequest, "dont have accookie")
// 	}
// 	rawtoken := cookie.Value
// 	if rawtoken == "" {
// 		return c.JSON(http.StatusUnauthorized, "null AC cookie")
// 	}

// 	///auth token
// 	_, err = jwt.ParseWithClaims(rawtoken, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("key")), nil
// 	})
// 	if err != nil {
// 		if err == jwt.ErrSignatureInvalid {
// 			return c.JSON(http.StatusUnauthorized, "signature is invalid")
// 		}
// 		return c.JSON(http.StatusForbidden, "Authorization failed")
// 	}
// 	return c.JSON(http.StatusOK, "done")
// }
