package main

import (
	"log"

	pb "github.com/louislef299/bash/projects/side_stuff/grpc-up-and-running/productinfo/ecommerce"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)
}
