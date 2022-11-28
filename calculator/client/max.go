package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/rugggger/go-grpc/calculator/proto"
)

func doMax(c pb.CalculatorServiceClient) {
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Could not sum %v", err)
	}

	go func() {
		stream.Send(&pb.MaxRequest{Num: 5})
		time.Sleep(2 * time.Second)
		stream.Send(&pb.MaxRequest{Num: 95})
		time.Sleep(2 * time.Second)
		stream.Send(&pb.MaxRequest{Num: 1})
		time.Sleep(2 * time.Second)
		stream.Send(&pb.MaxRequest{Num: 512})
		stream.CloseSend()
	}()

	endChan := make(chan string)
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				stream.CloseSend()
				endChan <- "end"
				break
			}
			if err != nil {
				log.Fatalf("Could not get factor %v", err)
			}
			fmt.Printf("Max received: %d\n", res.Num)
		}

	}()

	<-endChan
	close(endChan)
}
