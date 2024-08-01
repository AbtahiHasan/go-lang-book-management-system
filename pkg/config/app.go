package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var (
	db *gorm.DB
)

func ConnectDb() {
	godotenv.Load(".env")
	dbUrl := os.Getenv("DB_URL")
	d, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDb() *gorm.DB {
	return db
}