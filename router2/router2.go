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
	modify := e.Group("/modify")

	//////middleware//////
	modify.Use(middlewares.TokenchekMiddleware)

	w.Use(middlewares.TokenchekMiddleware)
	g.Use(middlewares.TokenchekMiddleware)
	////////////////////////////////////////

	modify.GET("/?id=", handler.UpdateBoard)
	modify.POST("/?id=", handler.Postupdate)
	e.GET("/view/?id=", handler.Readboard)
	e.GET("/ckecktoken", handler.Checktoken)
	e.GET("/api/signin", handler.SignIn)
	e.GET("/api/signout", handler.SignOut)

	w.GET("/write", handler.Boardform)   //to write page
	w.POST("/write", handler.WriteBoard) //upload wrote board
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
