package main

import (
	"context"
	"fmt"
	pb "github.com/shniu/gostuff/projects/gmq/internal/protocol/protobuf"
	"google.golang.org/grpc"
	"os"
	"sync"
	"time"
)

func main() {
	addr := "localhost:4321"
	fmt.Printf("Connecting to gmq server: %s\n", addr)

	// 这里的 addr 可能是 namesrv 的地址
	producer, err := NewProducer(addr)
	if err != nil || producer == nil {
		fmt.Println("Fail to start gmq.")
		os.Exit(-1)
	}

	err = producer.Send("mytopic-1", []byte("Hello, gmq"))
	if err != nil {
		fmt.Println("Failure to send message to mytopic-1.")
	}
}

func NewProducer(addr string) (*Producer, error) {
	p := &Producer{
		id:   1,
		addr: addr,
	}

	return p, nil
}

type Producer struct {
	id   int64
	addr string

	wg sync.WaitGroup
}

func (p *Producer) Send(topic string, data []byte) error {
	// len := len(data)

	payload := pb.Payload{
		Topic: topic,
		Len:   uint32(len(data)),
		Data:  data,
	}

	conn, err := grpc.Dial(p.addr, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	defer conn.Close()

	c := pb.NewMessageServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Send(ctx, &payload)
	if err != nil {
		fmt.Printf("could not send: %v", err)
		return err
	}
	fmt.Printf("Send result: %v \n", r)

	return nil
}
