package main

import (
	"fmt"

	pb "github.com/rugggger/go-grpc/calculator/proto"
)

func (s *Server) Factor(in *pb.FactorRequest, stream pb.CalculatorService_FactorServer) error {
	N := in.Number
	k := int64(2)
	for N/k != 1 {
		fmt.Println("try ", N, k)
		if N%k == 0 {
			fmt.Println("send", k)
			stream.Send(&pb.FactorResponse{Factor: int32(k)})
			N = N / k
			k = 2
		} else {
			k++
		}
	}
	fmt.Println("send", N)
	stream.Send(&pb.FactorResponse{Factor: int32(N)})

	return nil
}
