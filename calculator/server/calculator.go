package main

import (
	"context"
	"log"

	pb "github.com/rugggger/go-grpc/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Greet function was envoked with %v", in)
	return &pb.SumResponse{
		Result: int64(in.A + in.B),
	}, nil

}
