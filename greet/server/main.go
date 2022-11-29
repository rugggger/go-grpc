package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/rugggger/go-grpc/greet/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening in %s\n", addr)
	opts := []grpc.ServerOption{}
	tls := true
	if tls {
		certfile := "ssl/server.crt"
		keyfile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certfile, keyfile)
		if err != nil {
			log.Fatalf("Can't load certificates %v", err)
		}
		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatal("Failed to serve")
	}
}
