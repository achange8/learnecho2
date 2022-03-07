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
	USER := os.Getenv("username")
	PASS := os.Getenv("password")
	PROTOCOL := "tcp(database-2.cnxsbvp2uyyo.ap-northeast-2.rds.amazonaws.com)"
	DBNAME := os.Getenv("DBNAME")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Asia%2FSeoul"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	return db
}
