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
	m := e.Group("/modify")
	m.Use(middlewares.TokenchekMiddleware)
	//m.use check same user
	w.Use(middlewares.TokenchekMiddleware)
	g.Use(middlewares.TokenchekMiddleware)
	e.GET("/view", handler.Readboard)
	e.GET("/ckecktoken", handler.Checktoken)
	e.GET("/api/signin", handler.SignIn)
	e.GET("/api/signout", handler.SignOut)

	w.GET("/write", handler.Boardform)
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

//ToDo : make borad list, update list, delete board
//TODO : log out
