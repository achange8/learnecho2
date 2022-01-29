package handler

import (
	"net/http"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/helper"
	"github.com/achange8/learnecho2/models"
	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {
	user := new(models.User)
	err := c.Bind(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}
	db := db.Connect()
	//find email
	result := db.Find(&user, "email=?", user.Email)
	if result.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "existing email",
		})
	}

	//find id existing
	id := db.Find(&user, "id=?", user.Id)

	if id.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "existing id",
		})
	}
	//create pw -> hash val
	hashpw, err := helper.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}
	user.Password = hashpw

	if err := db.Create(&user); err.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Failed SignUp",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Success sign up!",
	})

}
