package client

import (
	"context"

	pb "github.com/louislef299/go-scripts/projects/mlctl/api/v1"
	mdtlog "github.com/louislef299/go-scripts/projects/mlctl/internal/log"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type Client struct {
	Log *mdtlog.Logger
}

func (c *Client) Dial() {
	c.Log.Info.Printf("attempting to dial to address %v", viper.GetString("server.address"))
	conn, err := grpc.Dial(viper.GetString("server.address"), grpc.WithInsecure())
	if err != nil {
		c.Log.Error.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	stub := pb.NewRecordClient(conn)

	r, err := stub.SendMessage(context.Background(), &pb.Request{Name: "Louis"})
	if err != nil {
		c.Log.Error.Fatalf("Could not send message: %v", err)
	}
	c.Log.Info.Printf("Message response: %v", r.Msg)
}
