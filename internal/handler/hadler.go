package handler

import (
	"context"
	"libnet/internal/service/grpc/pb"
)

func Init() *Handler {
	return &Handler{}
}

type Handler struct {
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
