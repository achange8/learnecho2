package models

import (
	"time"

	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

//Id = BOARD.WRITER
type User struct {
	Id       string
	Email    string
	Password string
}

type JwtClaims struct {
	jwt.StandardClaims
}

//todo remake table

type BOARD struct {
	NUM       int `gorm:"primaryKey"`
	TITLE     string
	WRITER    string
	CONTENT   string
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	HiTCOUNT  int
}

type Refresh struct {
	Id       string
	Reftoken string
}
