package bookRepository

import (
	"database/sql"
	"fmt"
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
func (b BookRepository) GetBooks(db *sql.DB, book models.Book, books []models.Book) ([]models.Book, error) {
	rows, err := db.Query("select * from books")
	logFatal(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		logFatal(err)

		books = append(books, book)
	}

	// возвращает список всех книг в БД в таблице БД books
	return books, nil
}

// Метод книжного репозитория, возвращающий книгу с определенным id в таблице books
func (b BookRepository) GetBook(db *sql.DB, id int) (models.Book, error) {
	var book models.Book

	rows := db.QueryRow("select * from books where id=$1", id)

	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	if err != nil {
		if err == sql.ErrNoRows {
			return book, fmt.Errorf("книга c id %v не найдена", id)
		} else {
			logFatal(err)
		}
	}

	// возвращает найденную книгу по её ID
	return book, nil
}

// Метод книжного репозитория, добавляющим новую книгу в таблицу books
func (b BookRepository) AddBook(db *sql.DB, book models.Book) (models.Book, error) {
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;",
		book.Title, book.Author, book.Year).Scan(&book.ID)
	if err != nil {
		return book, fmt.Errorf("не удалось создать книгу по причине: %v", err)
	}

	// возвращает созданную книгу (вместе с её ID, полученным из БД)
	return book, nil
}

// Метод книжного репозитория, обновляющий существующую книгу по её id в таблице books
func (b BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	// используем метод получения книги по её id
	_, errNotFinded := b.GetBook(db, book.ID)
	// если есть ошибка, значит книга не найдена, возвращаем ошибку
	if errNotFinded != nil {
		return 0, errNotFinded
	}

	result, err := db.Exec("update books set title=$1, author=$2, year=$3 where id=$4 RETURNING id",
		&book.Title, &book.Author, &book.Year, &book.ID)
	logFatal(err)

	rowsUpdated, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("не удалось обновить книгу по причине: %v", err)
	}

	// возвращает кол-во обновленных строк таблицы
	return rowsUpdated, nil
}

// Метод книжного репозитория, удяляющий существующую книгу по её id в таблице books
func (b BookRepository) RemoveBook(db *sql.DB, id int) (int64, error) {
	result, err := db.Exec("delete from books where id = $1", id)
	logFatal(err)

	rowsDeleted, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("не удалось удалить книгу по причине: %v", err)
	}

	// возвращает кол-во удаленных строк таблицы
	return rowsDeleted, nil
}
