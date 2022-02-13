package router2

import (
	"github.com/achange8/learnecho2/handler"
	"github.com/achange8/learnecho2/middlewares"

	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	g := e.Group("/user")
	write := e.Group("/board")
	modify := e.Group("/modify")

	//////middleware//////
	modify.Use(middlewares.TokenchekMiddleware)

	write.Use(middlewares.TokenchekMiddleware)
	g.Use(middlewares.TokenchekMiddleware)
	////////////////////////////////////////
	e.GET("/", handler.Mainboard)
	modify.GET("/?id=", handler.UpdateBoard)
	modify.POST("/?id=", handler.Postupdate)
	e.GET("/view/?id=", handler.Readboard)
	e.GET("/api/signin", handler.SignIn)
	e.GET("/api/signout", handler.SignOut)
	write.GET("/write", handler.Boardform)   //to write page
	write.POST("/write", handler.WriteBoard) //upload wrote board
	e.POST("/api/signup", handler.Signup)
	return e
}
