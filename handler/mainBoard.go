package handler

import (
	"net/http"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/models"
	"github.com/labstack/echo"
)

//method : get
//localhost/
func Mainboard(c echo.Context) error {
	var list []models.BOARD
	db := db.Connect()
	result := db.Table("boards").Find(&list)
	if result.RowsAffected == 0 {
		return c.JSON(http.StatusOK, "no result")
	}
	return c.JSON(http.StatusOK, list)
}
