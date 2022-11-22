package main

import (
	"log"

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
	clientCalc := pb.NewCalculatorServiceClient(conn)
	doGreet(client)
	doCalc(clientCalc)

}
