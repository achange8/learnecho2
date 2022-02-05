package models

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

//Id = BOARD.WRITER
type User struct {
	gorm.Model
	Id       string `json:id`
	Email    string `json:email`
	Password string `-`
}

type JwtClaims struct {
	jwt.StandardClaims
}

//todo remake table

type BOARD struct {
	gorm.Model
	NUM      int    `json:NUM`
	TITLE    string `json:TITLE`
	WRITER   string `json:WRITER`
	CONTENT  string `json:CONTENT`
	DB_DATE  string
	HiTCOUNT int
}

type Refresh struct {
	Id       string
	Reftoken string
}
