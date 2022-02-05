package handler

import (
	"net/http"
	"strconv"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//board/modify/?id=~
//method : get
func updateBoard(c echo.Context) error {
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
	username := claims["jti"].(string)

	//parse id in url
	id := c.QueryParam("id")
	//change string--> int
	num, numerr := strconv.Atoi(id)
	if numerr != nil {
		return c.JSON(http.StatusBadRequest, "page not found")
	}
	//select * from boards where id = {num}
	db := db.Connect()
	board := new(models.BOARD)
	if err := c.Bind(board); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	result := db.Raw("SELECT * FROM boards WHERE NUM = ?", num).Scan(&board)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "no result")
	}
	if username != board.WRITER {
		return c.JSON(http.StatusUnauthorized, "Only the writer can modify it. ")
	}
	return c.JSON(http.StatusOK, board)
}

//method post
//board/modify/?id=~
func postupdate(c echo.Context) error {
	//parse id in url
	id := c.QueryParam("id")
	//change string--> int
	num, numerr := strconv.Atoi(id)
	if numerr != nil {
		return c.JSON(http.StatusBadRequest, "page not found")
	}
	//select * from boards where id = {num}
	db := db.Connect()
	board := new(models.BOARD)
	if err := c.Bind(board); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	db.Model(&board).Where("NUM = ?", num).Updates(models.BOARD{TITLE: board.TITLE, CONTENT: board.CONTENT})

	return c.JSON(http.StatusOK, board)

}
