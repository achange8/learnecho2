package models

import (
	"github.com/golang-jwt/jwt"
)

//Id = BOARD.WRITER
type User struct {
	Id       string `json:id`
	Email    string `json:email`
	Password string `-`
}

type JwtClaims struct {
	jwt.StandardClaims
}

type BOARD struct {
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
