package main

import (
	"log"
	"net/http"

	"github.com/abtahihasan/go-lang-book-management-system/pkg/routes"
	"github.com/gorilla/mux"
)


func main () {
	r := mux.NewRouter()
	routes.RegisterBooksStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000",r))
}