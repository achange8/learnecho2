package models

import "github.com/golang-jwt/jwt"

type User struct {
	Id       string `json:id`
	Email    string `json:email`
	Password string `json:password`
}

type JwtClaims struct {
	jwt.StandardClaims
}

type BOARD struct {
	Id      int    `json:ID`
	Title   string `json:TITLE`
	Writer  string
	Content string `json:CONTENT`
	DB_DATE string
}

type Refresh struct {
	id       string `json:id`
	reftoken string `json:reftoken`
}
