package server

import (
	"context"
	"fmt"
	"net"

	pb "github.com/louislef299/bash/projects/mlctl/api/v1"
	mdtlog "github.com/louislef299/bash/projects/mlctl/internal/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	users []*pb.Request
	Log   *mdtlog.Logger
	pb.UnimplementedRecordServer
}

func (s *Server) SendMessage(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	s.users = append(s.users, in)
	return &pb.Response{Msg: fmt.Sprintf("hello %s", in.Name)}, status.New(codes.OK, "").Err()
}

func (s *Server) GetUsers(*pb.Request, pb.Record_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}

func (s *Server) StartServer(port int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		s.Log.Error.Fatalf("failed to listen: %v", err)
	}
	stub := grpc.NewServer()
	pb.RegisterRecordServer(stub, &Server{})

	s.Log.Info.Printf("setting server address to %v in config file", lis.Addr().String())
	viper.Set("server.address", lis.Addr().String())
	viper.WriteConfig()

	s.Log.Info.Printf("Starting gRPC listener on port %d", port)
	if err := stub.Serve(lis); err != nil {
		s.Log.Error.Fatalf("failed to serve: %v", err)
	}
}
