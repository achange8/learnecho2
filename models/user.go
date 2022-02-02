package models

import (
	"github.com/golang-jwt/jwt"
)

type User struct {
	Id       string `json:id`
	Email    string `json:email`
	Password string `-`
}

type JwtClaims struct {
	jwt.StandardClaims
}

type BOARD struct {
	Num      int
	Title    string `json:TITLE`
	Writer   string `json:WRITER`
	Content  string `json:CONTENT`
	DB_DATE  string
	HiTCOUNT int
}

type Refresh struct {
	Id       string
	Reftoken string
}
