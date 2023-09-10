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

func (s GrpcService) FindBook(ctx context.Context, req *pb.AuthorRequest) (*pb.BookResponse, error) {
	return s.handler.FindBook(ctx, req)
}

func (s GrpcService) FindAuthor(ctx context.Context, req *pb.BookRequest) (*pb.AuthorResponse, error) {
	return s.handler.FindAuthor(ctx, req)
}
