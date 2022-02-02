package handler

import (
	"net/http"

	"github.com/achange8/learnecho2/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

//board/id=~/update
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
	board := new(models.BOARD)
	board.Writer = claims["jti"].(string)
	return nil
}
