package app

import (
	"fmt"
	"libnet/internal/adaptor"
	grpcService "libnet/internal/service/grpc"
	"libnet/internal/service/grpc/pb"
	"libnet/pkg/config"
	"libnet/pkg/db"
	"log"
	"net"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type Config interface {
	GetString(key string) string
	GetInt(key string) int
}

func Init(configPath string) (*Server, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Println(err)
		return nil, errors.WithStack(err)
	}

	db, err := db.Init(cfg)
	if err != nil {
		log.Println(err)
		return nil, errors.WithStack(err)
	}
	repo := adaptor.Init(db.Conn)

	return &Server{config: cfg, repo: repo}, nil
}

type Server struct {
	config Config
	repo   *adaptor.Repository
}

func (s *Server) Start() error {
	listener, err := net.Listen(
		s.config.GetString("app.network"),
		fmt.Sprintf("%s:%s", s.config.GetString("app.host"), s.config.GetString("app.port")),
	)

	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterLibraryServer(grpcServer, grpcService.Init(s.repo))
	err = grpcServer.Serve(listener)

	return err
}
