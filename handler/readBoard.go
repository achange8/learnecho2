package handler

import (
	"net/http"
	"strconv"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/labstack/echo"
)

//method : get
//localhost/view/?id=**
func Readboard(c echo.Context) error {
	//parse id in url
	id := c.QueryParam("id")
	println(id)
	//change string--> int
	num, numerr := strconv.Atoi(id)
	if numerr != nil {
		return c.JSON(http.StatusBadRequest, "page not found1")
	}
	//select * from boards where id = {num}
	db := db.Connect()
	board := new(models.BOARD)
	if err := c.Bind(board); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request2",
		})
	}
	result := db.Raw("SELECT * FROM boards WHERE NUM = ?", num).Scan(&board)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusNotFound, "no result1")
	}
	db.Model(board).Where("NUM = ?", board.NUM).Update("hi_tcount", board.HiTCOUNT+1)

	return c.JSON(http.StatusOK, board)
}
