package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/rugggger/go-grpc/calculator/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("Failed to listen on %v\n", err)
	}
	log.Printf("Listening in %s\n", addr)
	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve")
	}
}
