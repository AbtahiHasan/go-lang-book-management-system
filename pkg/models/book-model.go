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


func (b *Book) CreateBook () *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func  GetAllBooks() *[]Book {
	var books []Book
	db.Find(&books)
	return &books
}

func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID=?",id).Find(&book)
	return &book, db
}

func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}