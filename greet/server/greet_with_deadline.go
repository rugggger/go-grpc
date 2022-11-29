package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/rugggger/go-grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadLine(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was envoked with %v", in)
	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			fmt.Println("Deadline exceeded! cancelled")
			return nil, status.Error(codes.Canceled, "Cancelled due to timeout")
		}

		fmt.Println("Waiting...")
		time.Sleep(1 * time.Second)
	}
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil

}
