package handler

import (
	"net/http"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/labstack/echo"
)

func Mainboard(c echo.Context) error {
	var list []models.BOARD
	db := db.Connect().Config
	board := new(models.BOARD)
	if err := c.Bind(board); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	result := db.Find(&board)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusOK, "no result")
	}

	return c.JSON(http.StatusOK, board)

}
