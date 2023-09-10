package handler

import (
	"context"
	"libnet/internal/adaptor"
	"libnet/internal/service/grpc/pb"
)

func Init(repo *adaptor.Repository) *Handler {
	return &Handler{repo: repo}
}

type Handler struct {
	repo *adaptor.Repository
}

func (h Handler) FindBook(ctx context.Context, req *pb.AuthorRequest) (*pb.BookResponse, error) {
	books, err := h.repo.FindBook(ctx, req.LastName)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Book, 0)
	for _, book := range books {
		res = append(res, &pb.Book{
			Id:    book.Id,
			Title: book.Title,
		})
	}

	return &pb.BookResponse{
		Books: res,
	}, nil
}

func (h Handler) FindAuthor(ctx context.Context, req *pb.BookRequest) (*pb.AuthorResponse, error) {
	authors, err := h.repo.FindAuthor(ctx, req.Title)
	if err != nil {
		return nil, err
	}

	res := make([]*pb.Author, 0)
	for _, author := range authors {
		res = append(res, &pb.Author{
			Id:       author.Id,
			LastName: author.LastName,
		})
	}

	return &pb.AuthorResponse{
		Authors: res,
	}, nil
}
