package main

import (
	"log"
	"time"

	pb "github.com/rugggger/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

var addr string = "localhost:50051"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certfile := "ssl/ca.crt"

		creds, err := credentials.NewClientTLSFromFile(certfile, "")
		if err != nil {
			log.Fatalf("Can't dialup %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))

	}
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect %v\n", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	doGreet(client)
	// doGreetManyTimes(client)
	doGreet2(client, time.Second*3)
	doGreet2(client, time.Second*1)
	doGreet2(client, time.Second*5)

}
