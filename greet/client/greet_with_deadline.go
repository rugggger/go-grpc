package main

import (
	"context"
	"log"
	"time"

	pb "github.com/rugggger/go-grpc/greet/proto"
	"google.golang.org/grpc/status"
)

func doGreet2(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	res, err := c.GreetWithDeadLine(ctx, &pb.GreetRequest{FirstName: "Yaron"})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			log.Printf("gRPC error (%d) %s\n", e.Code(), e.Message())
			return
		} else {
			log.Fatalf("Error unknown: could not greet %v\n", err)
		}
	}
	log.Printf("Greetings2: %s", res.Result)
}
