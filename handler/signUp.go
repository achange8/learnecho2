package handler

import (
	"fmt"
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
	//find email
	c.JSON(200, user)
	userID := user.Id
	result := db.Find(user, "email=?", user.Email)
	c.JSON(200, user)
	fmt.Println(result)
	if result.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "existing email")
	}

	//find id existing
	id := db.Find(user, "id=?", userID)
	if id.RowsAffected != 0 {
		return c.JSON(http.StatusBadRequest, "existing id")
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
