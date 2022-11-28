package main

import (
	"context"
	"fmt"
	"math"

	pb "github.com/rugggger/go-grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(c context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	n := in.Num
	if n < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number"),
		)
	}

	return &pb.SqrtResponse{
		Result: int64(math.Sqrt(float64(n))),
	}, nil

}
