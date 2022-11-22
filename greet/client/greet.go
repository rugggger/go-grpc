package main

import (
	"context"
	"log"

	pb "github.com/rugggger/go-grpc/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Yaron"})
	if err != nil {
		log.Fatalf("Could not greet %v", err)
	}
	log.Printf("Greetings: %s", res.Result)
}
