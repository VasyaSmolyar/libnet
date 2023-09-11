package action

import (
	"context"
	"libnet/internal/model"
)

type FindBooks struct {
	repo BookRepository
}

func NewFindBooks(repo BookRepository) *FindBooks {
	return &FindBooks{repo: repo}
}

func (act FindBooks) Do(ctx context.Context, title string) ([]*model.Book, error) {
	return act.repo.FindBook(ctx, title)
}
