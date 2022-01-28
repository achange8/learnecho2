package router2

import (
	"go/token"
	"net/http"

	"github.com/achange8/learnecho2/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	g := e.Group("/user")
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey: []byte(os.Getenv("key")),
	}))
	g.POST("/profile", Uprofile)
	e.POST("/api/signin", handler.SignIn)

	e.POST("/api/signup", handler.SighUp)
	return e
}

func Uprofile(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	log.Println("User email :",claims["Id"])
	
	return c.String(http.StatusOK, "logged in user page")
}
func checklogin(next echo.HandlerFunc) echo.HandlerFunc{
	cookie,err := c.Cookie("JWTaccessCookie")
	if err != nil {
		return err
	}
	if cookie == ""
}
