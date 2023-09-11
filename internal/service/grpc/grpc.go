package grpc

import (
	"context"
	"libnet/internal/action"
	"libnet/internal/service/grpc/pb"
)

func Init(repo action.Repository) pb.LibraryServer {
	return &GrpcService{repo: repo}
}

type GrpcService struct {
	pb.UnimplementedLibraryServer
	repo action.Repository
}

func (s GrpcService) FindBook(ctx context.Context, req *pb.AuthorRequest) (*pb.BookResponse, error) {
	act := action.NewFindBooks(s.repo)

	books, err := act.Do(ctx, req.LastName)
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

func (s GrpcService) FindAuthor(ctx context.Context, req *pb.BookRequest) (*pb.AuthorResponse, error) {
	act := action.NewFindAuthors(s.repo)

	authors, err := act.Do(ctx, req.Title)
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
