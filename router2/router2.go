package router2

import (
	"net/http"

	"github.com/achange8/learnecho2/handler"
	"github.com/achange8/learnecho2/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	g := e.Group("/user")
	w := e.Group("/board")
	w.Use(middlewares.TokenchekMiddleware)
	middlewares.Setlogincheker(g)
	e.GET("/ckecktoken", handler.Checktoken)
	e.GET("/api/signin", handler.SignIn)
	w.GET("/write", handler.Boardform)
	g.GET("/profile", Uprofile)
	w.POST("/write", handler.WriteBoard) //upload board
	e.POST("/api/signup", handler.Signup)
	return e
}

func Uprofile(c echo.Context) error {
	// user := c.Get("user")
	// token := user.(*jwt.Token)
	// claims := token.Claims.(jwt.MapClaims)
	// log.Println("User ID: ", claims["jti"])

	return c.String(http.StatusOK, "logged in user page")
}
