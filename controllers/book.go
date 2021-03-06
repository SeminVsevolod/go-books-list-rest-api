package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"go-books-list-rest-api/models"
	bookRepo "go-books-list-rest-api/repository/book"
	"log"
	"net/http"
	"strconv"
)

type Controller struct{}

var books []models.Book

// вывести лог с ошибкой
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Метод контроллера, возвращающий все книги в формате json
func (c Controller) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		books = []models.Book{}
		bookRepo := bookRepo.BookRepository{}
		var err error
		books, err = bookRepo.GetBooks(db, book, books)

		// возвращаем код 400 и ошибку
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// возвращаем список книг, преобразованный в json
		json.NewEncoder(w).Encode(books)
	}
}

// Метод контроллера, возвращающий книгу в формате json по её id
func (c Controller) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		// получаем параметры запроса в виде словаря map[string]string
		params := mux.Vars(r)
		bookRepo := bookRepo.BookRepository{}

		// преобразовываем id из str в int
		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		// используем книжный репозиторий для вызова метода GetBook, возвращает книгу по её id
		book, err = bookRepo.GetBook(db, id)

		// возвращаем код 404 и ошибку если книга с этим id не найдена
		if err != nil {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// возвращаем книгу с нужным id, преобразованную в json
		json.NewEncoder(w).Encode(book)
	}
}

// Метод контроллера, создающий новую книгу (возвращает созданную книгу)
func (c Controller) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		bookRepo := bookRepo.BookRepository{}

		// преобразуем из json, переданного в Body в запросе, в экземпляр структуры книги book
		json.NewDecoder(r.Body).Decode(&book)

		// используем книжный репозиторий для вызова метода AddBook, который добавляет новую книгу в таблицу книг, и возвращает созданную книгу
		book, err := bookRepo.AddBook(db, book)

		// возвращаем код 400 и ошибку если не удалось создать книгу
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// возвращаем созданную книгу, преобразованную в json
		json.NewEncoder(w).Encode(book)
	}
}

// Метод контроллера, обновляющий книгу по ID (возвращает кол-во обновленных строк БД)
func (c Controller) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		// преобразуем из json, переданного в Body в запросе, в экземпляр структуры книги book
		json.NewDecoder(r.Body).Decode(&book)
		bookRepo := bookRepo.BookRepository{}

		// используем книжный репозиторий для вызова метода UpdateBook, который обновляет существующую книгу по её id, и возвращает кол-во обновленных строк
		rowsUpdated, err := bookRepo.UpdateBook(db, book)

		// возвращаем код 400 и ошибку если не удалось обновить книгу
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// возвращаем кол-во обновленных строк, преобразованное в json
		json.NewEncoder(w).Encode(map[string]int64{"rowsUpdated": rowsUpdated})
	}
}

// Метод контроллера, удаляющий книгу по её ID (возвращает кол-во удаленных строк БД)
func (c Controller) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// получаем параметры запроса в виде словаря map[string]string
		params := mux.Vars(r)
		bookRepo := bookRepo.BookRepository{}

		// преобразовываем id из str в int
		id, err := strconv.Atoi(params["id"])
		logFatal(err)

		// используем книжный репозиторий для вызова метода RemoveBook, который удаляет книгу по её id, и возвращает кол-во удаленных строк
		rowsDeleted, err := bookRepo.RemoveBook(db, id)

		// возвращаем код 400 и ошибку если не удалось удалить книгу
		if err != nil {
			w.WriteHeader(400)
			json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
			return
		}

		// возвращаем кол-во удаленных строк, преобразованное в json
		json.NewEncoder(w).Encode(map[string]int64{"rowsDeleted": rowsDeleted})
	}
}
