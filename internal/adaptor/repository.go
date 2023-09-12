package adaptor

import (
	"context"
	"database/sql"
	"libnet/internal/model"

	"github.com/pkg/errors"
)

func Init(db *sql.DB) *Repository {
	return &Repository{db: db}
}

type Repository struct {
	db *sql.DB
}

func (repo Repository) FindBooks(ctx context.Context, lastName string) ([]*model.Book, error) {
	books := make([]*model.Book, 0)

	var query = "SELECT books.id, books.title FROM books JOIN book_author ON books.id = book_author.book_id JOIN authors ON authors.id = book_author.author_id WHERE authors.last_name = ?;"
	rows, err := repo.db.QueryContext(ctx, query, lastName)
	if err != nil {
		return books, errors.WithStack(err)
	}

	for rows.Next() {
		var id, title string
		if err := rows.Scan(&id, &title); err != nil {
			return books, errors.WithStack(err)
		}
		books = append(books, &model.Book{
			Id:    id,
			Title: title,
		})
	}

	return books, errors.WithStack(err)
}

func (repo Repository) FindAuthors(ctx context.Context, title string) ([]*model.Author, error) {
	authors := make([]*model.Author, 0)

	var query = "SELECT authors.id, authors.last_name FROM authors JOIN book_author ON authors.id = book_author.author_id JOIN books ON books.id = book_author.book_id WHERE books.title = ?;"
	rows, err := repo.db.QueryContext(ctx, query, title)
	if err != nil {
		return authors, errors.WithStack(err)
	}

	for rows.Next() {
		var id, lastName string
		if err := rows.Scan(&id, &lastName); err != nil {
			return authors, errors.WithStack(err)
		}
		authors = append(authors, &model.Author{
			Id:       id,
			LastName: lastName,
		})
	}

	return authors, errors.WithStack(err)
}
