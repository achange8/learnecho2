package handler

import (
	"net/http"
	"strconv"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//method : delete
//must check access token
//soft delete
func DeleteBoard(c echo.Context) error {

	//check user
	cookie, err := c.Cookie("JWTaccessToken")
	if err != nil {
		return err
	}
	rawtoken := cookie.Value
	token, err := jwt.Parse(rawtoken, nil)
	if token == nil {
		return err
	}
	claims, _ := token.Claims.(jwt.MapClaims)
	username := claims["jti"].(string)
	id := c.QueryParam("id")
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
		return c.JSON(http.StatusUnauthorized, "only writer can delete board!")
	}
	// Batch Delete
	db.Where("NUM = ?", num).Delete(&board)
	return c.JSON(http.StatusOK, "delete done!")
	// UPDATE boards SET deleted_at="2013-10-29 10:23" WHERE NUM = num;
}
