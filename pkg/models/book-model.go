package models

import (
	"github.com/abtahihasan/go-lang-book-management-system/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title  string	`json:"title"`
	Author string	`json:"author"`
	Publication string	`json:"publication"`
}


func init () {
	config.ConnectDb()
	db = config.GetDb()
	db.AutoMigrate(&Book{})
}