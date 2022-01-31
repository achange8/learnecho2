package handler

import (
	"net/http"

	"github.com/achange8/learnecho2/db"
	"github.com/achange8/learnecho2/helper"
	"github.com/achange8/learnecho2/models"
	"github.com/labstack/echo"
)

func Signup(c echo.Context) error {
	db := db.Connect()
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "bad request",
		})
	}

	savepassword := user.Password
	//find id,email existing
	id := db.Find(user, "id=?", user.Id)
	result := db.Raw("SELECT * FROM users WHERE email = ?", user.Email).Scan(&user)
	if id.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "existing id")
	}
	if result.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "existing email")
	}

	//create pw -> hash val
	hashpw, err := helper.HashPassword(savepassword)
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
