package main

import (
	"log"
	"time"

	pb "github.com/rugggger/go-grpc/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
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
