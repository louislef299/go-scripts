package main

import (
	"context"
	"log"
	"time"

	pb "github.com/louislef299/bash/projects/side_stuff/grpc-up-and-running/productinfo/ecommerce"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewProductInfoClient(conn)

	name := "Apple iPhone 11"
	description := `Meet Apple iPhone 11. All-new dual-camera system with
	Ultra Wide and Night mode.`
	price := float32(1000.0)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AddProduct(ctx, &pb.Product{Name: name, Description: description, Price: price})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Product ID: %s added successfully", r.Id)

	product, err := c.GetProduct(ctx, &pb.ProductID{Id: r.Id})
	if err != nil {
		log.Fatalf("Could not get product: %v", err)
	}
	log.Printf("Product: %s", product.String())
}
