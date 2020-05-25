package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"
	"go-books-list-rest-api/controllers"
	"go-books-list-rest-api/database"
	"go-books-list-rest-api/models"
	"log"
	"net/http"
	_ "strconv"
)

var books []models.Book
var db *sql.DB

func init() {
	gotenv.Load()
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = database.ConnectDB()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/book/{id}", controller.GetBook(db)).Methods("GET")
	router.HandleFunc("/book", controller.AddBook(db)).Methods("POST")
	router.HandleFunc("/book", controller.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/book/{id}", controller.RemoveBook(db)).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
