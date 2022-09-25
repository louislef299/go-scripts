package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	pb "github.com/louislef299/bash/projects/coffee_shop_playgound/chat-grpc/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// block sigint
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	defer close(sigs)
	defer close(done)
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	conn, err := grpc.Dial("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// create client listener
	msgs := make(chan pb.Response)
	defer close(msgs)
	go func() {
		for {
			var resp pb.Response
			err = r.RecvMsg(&resp)
			if err == io.EOF {
				break
			}
			msgs <- resp
		}
		<-done
	}()

	fmt.Println("Enter your name:")
	var name string
	fmt.Scanln(&name)
	if err := r.Send(&pb.Request{Name: name}); err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	select {
	case <-done:
		r.CloseSend()
		log.Fatal("got the done signal")
	case msg, val := <-msgs:
		fmt.Println("got message:", msg.GetMsg(), val)
	}
}
