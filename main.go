package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	ID		int 	`json:id`
	Title 	string 	`json:title`
	Author 	string 	`json:author`
	Year 	string 	`json:year`
}

var books []Book

func main()  {
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", addBook).Methods("POST")
	router.HandleFunc("/book", updateBook).Methods("PUT")
	router.HandleFunc("/book{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets all books")
}

func getBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Gets one book")
}

func addBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Adds one book")
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Updates one book")
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	log.Println("Removes one book")
}
