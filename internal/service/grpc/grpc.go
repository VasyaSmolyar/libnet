package grpc

import (
	"context"
	"libnet/internal/handler"
	"libnet/internal/service/grpc/pb"
)

func Init(handler *handler.Handler) pb.LibraryServer {
	return &GrpcService{handler: handler}
}

type GrpcService struct {
	pb.UnimplementedLibraryServer
	handler *handler.Handler
}

func (s GrpcService) FindBook(context.Context, *pb.AuthorRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{
		Title: "Book",
	}, nil
}

func (s GrpcService) FindAuthor(context.Context, *pb.BookRequest) (*pb.AuthorResponse, error) {
	return &pb.AuthorResponse{
		LastName: "Smith",
	}, nil
}
