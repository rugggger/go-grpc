package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/rugggger/go-grpc/calculator/proto"
)

func doFactor(c pb.CalculatorServiceClient) {
	stream, err := c.Factor(context.Background(), &pb.FactorRequest{Number: 18})
	if err != nil {
		log.Fatalf("Could not sum %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not get factor %v", err)
		}
		fmt.Printf("got factor %d\n", res.Factor)
	}
}
