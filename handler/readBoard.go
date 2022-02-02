package handler

import (
	"net/http"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/labstack/echo"
)

//method : get board num
//board/view/?num=**
func Readboard(c echo.Context) error {
	//select * from boards where id = {num}
	db := db.Connect()
	board := new(models.BOARD)
	if err := c.Bind(board); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	err := db.Raw("SELECT * FROM boards WHERE NUM = ?", board.Num).Scan(board)
	if err.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, err)
	}
	return c.JSON(http.StatusOK, board)
}
