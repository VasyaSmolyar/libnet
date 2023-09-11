package action

import (
	"context"
	"libnet/internal/model"
)

type FindAuthors struct {
	repo AuthorRepository
}

func NewFindAuthors(repo AuthorRepository) *FindAuthors {
	return &FindAuthors{repo: repo}
}

func (act FindAuthors) Do(ctx context.Context, title string) ([]*model.Author, error) {
	return act.repo.FindAuthor(ctx, title)
}
