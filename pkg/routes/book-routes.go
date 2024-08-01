package routes

import (
	"github.com/abtahihasan/go-lang-book-management-system/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBooksStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/books/:id", controllers.GetBook).Methods("GET")
	router.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/books/:id", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/books/:id", controllers.DeleteBook).Methods("DELETE")
}