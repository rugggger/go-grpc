package main

import (
	"context"
	"io"
	"log"

	pb "github.com/rugggger/go-grpc/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {

	req := &pb.GreetRequest{FirstName: "Yaron"}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not greet %v", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream %v", err)
		}
		log.Printf("Greet %s", msg.Result)
	}
}
