package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:id`
	Title  string `json:title`
	Author string `json:author`
	Year   string `json:year`
}

var books []Book

func main() {
	router := mux.NewRouter()

	books = append(books, Book{ID: 1, Title: "The Go Programming Language", Author: "Alan A. A. Donovan, Brian W. Kernighan", Year: "2015"},
		Book{ID: 2, Title: "Concurrency in Go: Tools and Techniques for Developers", Author: "Cox-Buday, Katherine", Year: "2017"},
		Book{ID: 3, Title: "Go in Action", Author: "William Kennedy, Brian Ketelsen, Erik St. Martin", Year: "2015"},
		Book{ID: 4, Title: "An Introduction to Programming in Go", Author: "Caleb Doxsey", Year: "2012"},
		Book{ID: 5, Title: "Introducing Go: Build Reliable, Scalable Programs", Author: "Doxsey, Caleb", Year: "2016"})

	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/book/{id}", getBook).Methods("GET")
	router.HandleFunc("/book", addBook).Methods("POST")
	router.HandleFunc("/book", updateBook).Methods("PUT")
	router.HandleFunc("/book/{id}", removeBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	numericID, _ := strconv.Atoi(params["id"])

	for _, book := range books {
		if book.ID == numericID {
			json.NewEncoder(w).Encode(&book)
		}
	}
}

func addBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	books = append(books, book)

	json.NewEncoder(w).Encode(books)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)

	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}

	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	numericID, _ := strconv.Atoi(params["id"])

	for i, item := range books {
		if item.ID == numericID {
			books = append(books[:i], books[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(books)
}
