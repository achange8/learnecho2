package db

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {

	envERR := godotenv.Load(".env")
	if envERR != nil {
		fmt.Println("Could not load .env file")
		os.Exit(1)
	}
	USER := os.Getenv("DBUSER")
	PASS := os.Getenv("DBPASS")
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := os.Getenv("DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
