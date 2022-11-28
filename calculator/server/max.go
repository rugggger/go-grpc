package main

import (
	"fmt"
	"io"

	pb "github.com/rugggger/go-grpc/calculator/proto"
)

func (s *Server) Max(stream pb.CalculatorService_MaxServer) error {
	var max int32

	var n int32

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Close stream")
			return nil
		}
		if err != nil {
			fmt.Printf("Error %v", err)
		}
		n = res.Num
		if n > max {
			stream.Send(&pb.MaxResponse{Num: n})
			max = n

		}

	}

}
