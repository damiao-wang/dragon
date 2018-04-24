package main

import (
	"context"
	"log"
	"time"

	"wangjh/myapi2/rpc/pb"

	"google.golang.org/grpc"
)

const (
	Addr = "localhost:1400"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Sum(ctx, &pb.SumReq{
		A: 5,
		B: 6,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Printf("Resp: %v", r.V)
}
