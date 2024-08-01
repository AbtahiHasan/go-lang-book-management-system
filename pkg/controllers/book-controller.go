package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abtahihasan/go-lang-book-management-system/pkg/models"
	"github.com/gorilla/mux"
)

var books models.Book

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

