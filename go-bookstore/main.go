package main

import (
	"log"
	"net/http"

	"go-bookstore/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/search", handlers.SearchBooks).Methods("GET")
	r.HandleFunc("/books", handlers.GetBooks).Methods("GET")
	r.HandleFunc("/books", handlers.CreateBook).Methods("POST")
	r.HandleFunc("/books/{id}", handlers.GetBookByID).Methods("GET")
	r.HandleFunc("/books/{id}", handlers.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", handlers.DeleteBook).Methods("DELETE")
	r.HandleFunc("/books/search", handlers.SearchBooks).Methods("GET")

	log.Println("Server started on port 8080")
	http.ListenAndServe(":8080", r)
}
