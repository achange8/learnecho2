package handler

import (
	"net/http"
	"time"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//get html,css form to write
func Boardform(e echo.Context) error {
	return e.JSON(http.StatusOK, "this is html form to write board only signin users")
}

//post board ~/write
func WriteBoard(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	board := new(models.BOARD)
	board.Writer = claims["jti"].(string)
	board.DB_DATE = time.Now().Format("15:04")
	err := c.Bind(board)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	db := db.Connect()
	if err := db.Create(&board); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, "failed write!")
	}
	return c.JSON(http.StatusOK, "done!")
}

// func updateBoard(c echo.Context) error {
// 	//load data from db to update

// 	board := new(models.BOARD)
// 	err := c.Bind(board)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, "bad request")
// 	}
// 	db := db.Connect()
// 	if err := db.Create(&board); err.Error != nil {
// 		return c.JSON(http.StatusInternalServerError, "failed write!")
// 	}
// 	return c.JSON(http.StatusOK, "done!")
// }
