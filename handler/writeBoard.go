package handler

import (
	"net/http"
	"time"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//GET html,css form to write
//~/borad/list
func Boardform(e echo.Context) error {
	return e.JSON(http.StatusOK, "this is html form to write board only signin users")
}

//POST ~/board/write
func WriteBoard(c echo.Context) error {
	cookie, err := c.Cookie("JWTaccessToken")
	if err != nil { //TODO : reqeust refresh token ,create new actoken or login again
		return c.JSON(http.StatusBadRequest, "You dont have accookie")
	}

	//get claims from jwt token without validation
	rawtoken := cookie.Value
	token, err := jwt.Parse(rawtoken, nil)
	if token == nil {
		return err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	board := new(models.BOARD)
	board.Writer = claims["jti"].(string)
	board.DB_DATE = time.Now().Format("2006-01-02 15:04") //format time
	err = c.Bind(board)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	db := db.Connect()
	if err := db.Create(&board); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, "failed write!")
	}
	return c.JSON(http.StatusOK, "done!")
}
