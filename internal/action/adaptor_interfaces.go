package action

import (
	"context"
	"libnet/internal/model"
)

//go:generate mkdir -p dimock
//go:generate mockgen -source=adaptor_interfaces.go -destination=dimock/adaptor_interfaces_mock.go -package=dimock

type BookRepository interface {
	FindBooks(ctx context.Context, lastName string) ([]*model.Book, error)
}

type AuthorRepository interface {
	FindAuthors(ctx context.Context, title string) ([]*model.Author, error)
}

type Repository interface {
	BookRepository
	AuthorRepository
}
