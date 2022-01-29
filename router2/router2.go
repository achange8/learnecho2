package router2

import (
	"log"
	"net/http"

	"github.com/achange8/learnecho2/handler"
	"github.com/achange8/learnecho2/middlewares"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	g := e.Group("/user")
	w := e.Group("/board")
	middlewares.Setlogincheker(w)
	middlewares.Setlogincheker(g)
	w.GET("/write", handler.Boardform)
	w.POST("/write", handler.WriteBoard) //upload board
	g.GET("/profile", Uprofile)
	e.POST("/api/signin", handler.SignIn)
	e.POST("/api/signup", handler.Signup)
	return e
}

func Uprofile(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	log.Println("User ID: ", claims["jti"])

	return c.String(http.StatusOK, "logged in user page")
}
