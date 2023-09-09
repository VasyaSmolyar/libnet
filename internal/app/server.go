package app

import (
	"fmt"
	"libnet/internal/handler"
	grpcService "libnet/internal/service/grpc"
	"libnet/internal/service/grpc/pb"
	"libnet/pkg/config"
	"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func Init(configPath string) (*Server, error) {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &Server{config: cfg, Handler: handler.Init()}, nil
}

type Server struct {
	config  *viper.Viper
	Handler *handler.Handler
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

	pb.RegisterLibraryServer(grpcServer, grpcService.Init(s.Handler))
	err = grpcServer.Serve(listener)

	return err
}
