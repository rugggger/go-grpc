package main

import (
	"context"
	"log"

	pb "github.com/rugggger/go-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was envoked with %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil

}

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was envoked with %v", in)
	return &pb.SumResponse{
		Result: int64(in.A + in.B),
	}, nil

}
