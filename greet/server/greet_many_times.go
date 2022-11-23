package main

import (
	pb "github.com/rugggger/go-grpc/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {

	for i := 1; i < 10; i++ {
		stream.Send(&pb.GreetResponse{Result: "Hi" + in.FirstName})
	}

	return nil

}
