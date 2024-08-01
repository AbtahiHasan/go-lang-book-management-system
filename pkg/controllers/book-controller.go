package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abtahihasan/go-lang-book-management-system/pkg/models"
	"github.com/abtahihasan/go-lang-book-management-system/pkg/utils"
	"github.com/gorilla/mux"
)



func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	w.Header().Add("Content-Type", "application/json")
	res,_ := json.Marshal(books)

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(Id)
	res,_ := json.Marshal(bookDetails)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	Book := &models.Book{}
	utils.ParseBody(r, Book)
	b := Book.CreateBook()
	res,_ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook (w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	Id, err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := &models.Book{}
	utils.ParseBody(r, book)
	bookDetails,db := models.GetBookById(Id)
	if book.Title != "" {
		bookDetails.Title = book.Title
	}
	if book.Author != "" {
		bookDetails.Title = book.Author
	}
	if book.Publication != "" {
		bookDetails.Title = book.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}