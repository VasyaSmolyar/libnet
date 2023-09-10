package handler

import (
	"libnet/internal/adaptor"
	"libnet/internal/service/grpc/pb"
)

func Init(repo *adaptor.Repository) *Handler {
	return &Handler{repo: repo}
}

type Handler struct {
	repo *adaptor.Repository
}

func (h Handler) FindBook(req *pb.AuthorRequest) (*pb.BookResponse, error) {
	books, err := h.repo.FindBook(req.LastName)
	if err != nil {
		return nil, err
	}

	return &pb.BookResponse{
		Title: books[0].Title,
	}, nil
}

func (h Handler) FindAuthor(req *pb.BookRequest) (*pb.AuthorResponse, error) {
	authors, err := h.repo.FindAuthor(req.Title)
	if err != nil {
		return nil, err
	}

	return &pb.AuthorResponse{
		LastName: authors[0].LastName,
	}, nil
}
