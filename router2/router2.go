package router2

import (
	"github.com/achange8/learnecho2/handler"
	"github.com/labstack/echo"
)

func New() *echo.Echo {
	e := echo.New()
	e.POST("/api/signin", handler.SignIn)
	e.POST("/api/signup", handler.SighUp)
	return e
}
