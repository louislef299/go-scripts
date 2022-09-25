package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/louislef299/bash/projects/coffee_shop_playgound/chat-grpc/chat"
	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) HandleRequest(stream pb.Greeter_SayHelloServer) error {
	for {
		name, err := stream.Recv()
		fmt.Println("received name:", name)
		if err == io.EOF { // close the stream
			stream.Send(&pb.Response{Msg: "session closed"})
			return nil
		}
		if err != nil {
			return err
		}

		stream.Send(&pb.Response{Msg: fmt.Sprintf("Hello %v you fucker ;)", name)})
	}
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 5051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	log.Fatalf("failed to serve: %v", s.Serve(lis))
}
