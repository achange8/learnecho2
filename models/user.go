package models

import "github.com/golang-jwt/jwt"

type User struct {
	Id       int    `json:id`
	Email    string `json:email`
	Password string `json:password`
}

type JwtClaims struct {
	jwt.StandardClaims
}
