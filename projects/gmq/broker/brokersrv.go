package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	pb "github.com/shniu/gostuff/projects/gmq/internal/protocol/protobuf"
	"os"
)

type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) Send(ctx context.Context, payload *pb.Payload) (*pb.Reply, error) {
	fmt.Printf("Received: topic is %s, data is %s \n", payload.GetTopic(), payload.GetData())

	return &pb.Reply{
		Code: "0",
	}, nil
}

func main() {
	fmt.Println("MQ Broker...")

	l, err := net.Listen("tcp", "localhost:4321")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(-1)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server{})
	fmt.Printf("listening at %v \n", l.Addr())

	if err = s.Serve(l); err != nil {
		fmt.Printf("failed to start server %v", l.Addr())
	}
}
