package bookRepository

import (
	"database/sql"
	"go-books-list-rest-api/models"
	"log"
)

type BookRepository struct{}

// вывести лог с ошибкой
func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Метод книжного репозитория, возвращающий все книги в БД в таблице books
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) []models.Book {
	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	// возвращает список всех книг в БД в таблице БД books
	return books
}

// Метод книжного репозитория, возвращающий книгу с определенным id в таблице books
func (b BookRepository) GetBook(db *sql.DB, book models.Book, id int) models.Book {
	rows := db.QueryRow("select * from books where id=$1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
	logFatal(err)

	// возвращает найденную книгу по её ID
	return book
}

// Метод книжного репозитория, добавляющим новую книгу в таблицу books
func (b BookRepository) AddBook(db *sql.DB, book models.Book) models.Book {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&book.ID)

	logFatal(err)

	// возвращает созданную книгу (вместе с её ID, полученным из БД)
	return book
}

// Метод книжного репозитория, обновляющий существующую книгу по её id в таблице books
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) int64 {
	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)

	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	logFatal(err)

	// возвращает кол-во обновленных строк таблицы
	return rowsUpdated
}

// Метод книжного репозитория, удяляющий существующую книгу по её id в таблице books
func (b BookRepository) RemoveBook(db *sql.DB, id int) int64 {
	result, err := db.Exec("delete from books where id = $1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	logFatal(err)

	// возвращает кол-во удаленных строк таблицы
	return rowsDeleted
}
