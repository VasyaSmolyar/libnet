package adaptor

import (
	"database/sql"
	"libnet/internal/model"
)

func Init(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sql.DB
}

func (repo Repository) FindBook(lastName string) ([]*model.Book, error) {
	books := make([]*model.Book, 0)

	var query = "SELECT books.id, books.title FROM books JOIN book_author ON books.id = book_author.book_id JOIN authors ON authors.id = book_author.author_id WHERE authors.last_name = $1;"
	rows, err := repo.db.Query(query)
	for rows.Next() {
		var id, title string
		if err := rows.Scan(&id, &title); err != nil {
			return books, err
		}
		books = append(books, &model.Book{
			Id:    id,
			Title: title,
		})
	}

	return books, err
}

func (repo Repository) FindAuthor(title string) ([]*model.Author, error) {
	authors := make([]*model.Author, 0)

	var query = "SELECT authors.last_name FROM authors JOIN book_author ON authors.id = book_author.author_id JOIN books ON books.id = book_author.book_id WHERE books.title = $1;"
	rows, err := repo.db.Query(query)
	for rows.Next() {
		var id, lastName string
		if err := rows.Scan(&id, &title); err != nil {
			return authors, err
		}
		authors = append(authors, &model.Author{
			Id:       id,
			LastName: lastName,
		})
	}

	return authors, err
}
