package action

import (
	"context"
	"libnet/internal/model"
)

//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g
type BookRepository interface {
	FindBook(ctx context.Context, lastName string) ([]*model.Book, error)
}

type AuthorRepository interface {
	FindAuthor(ctx context.Context, title string) ([]*model.Author, error)
}

type Repository interface {
	BookRepository
	AuthorRepository
}
