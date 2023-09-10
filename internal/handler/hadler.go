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

func (h Handler) FindBook(context.Context, *pb.AuthorRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{
		Title: "Book",
	}, nil
}

func (s Handler) FindAuthor(context.Context, *pb.BookRequest) (*pb.AuthorResponse, error) {
	return &pb.AuthorResponse{
		LastName: "Smith",
	}, nil
}
