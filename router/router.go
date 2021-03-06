package router

import (
	"net/http"

	"github.com/achange8/learnecho2/handler"
	"github.com/achange8/learnecho2/middlewares"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	g := e.Group("/user")
	write := e.Group("/board")
	modify := e.Group("/modify")

	//////middleware//////
	modify.Use(middlewares.TokenchekMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	})) //CORS
	write.Use(middlewares.TokenchekMiddleware)
	g.Use(middlewares.TokenchekMiddleware)
	////////////////////////////////////////
	g.GET("/check", handler.Usercheck)       // done
	e.GET("/", handler.Mainboard)            // done
	modify.POST("/", handler.Postupdate)     //done
	modify.DELETE("/", handler.DeleteBoard)  //done
	e.GET("/view/", handler.Readboard)       // done
	e.POST("/api/signup", handler.Signup)    // done
	e.POST("/api/signin", handler.SignIn)    // done
	e.GET("/api/signout", handler.SignOut)   // done
	write.POST("/write", handler.WriteBoard) //upload wrote board
	return e
}
