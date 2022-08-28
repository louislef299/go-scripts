package server

import (
	"context"
	"fmt"

	pb "github.com/louislef299/bash/projects/mlctl/api/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	users []*pb.Request
	pb.UnimplementedRecordServer
}

func (s *Server) SendMessage(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	s.users = append(s.users, in)
	return &pb.Response{Msg: fmt.Sprintf("hello %s", in.Name)}, status.New(codes.OK, "").Err()
}

func (s *Server) GetUsers(*pb.Request, pb.Record_GetUsersServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
