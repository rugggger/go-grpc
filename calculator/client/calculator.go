package main

import (
	"context"
	"log"

	pb "github.com/rugggger/go-grpc/calculator/proto"
)

func doCalc(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{A: 10, B: 3})
	if err != nil {
		log.Fatalf("Could not sum %v", err)
	}
	log.Printf("Sum: %s", res.Result)
}
