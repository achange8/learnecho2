package middlewares

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Setlogincheker(g *echo.Group) {
	envERR := godotenv.Load(".env")
	if envERR != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
	}

	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("key")),
		TokenLookup:   "cookie:JWTaccessCookie",
	}))
}

func SetRefreshChecker(g *echo.Group) {
	envERR := godotenv.Load(".env")
	if envERR != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
	}
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte(os.Getenv("key2")),
		TokenLookup:   "cookie:JWTRefreshToken",
	}))
}
