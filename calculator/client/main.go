package main

import (
	"log"

	pb "github.com/rugggger/go-grpc/calculator/proto"
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

	clientCalc := pb.NewCalculatorServiceClient(conn)
	doCalc(clientCalc)
	doFactor(clientCalc)
	doMax(clientCalc)
	doSqrt(clientCalc, 243)
	doSqrt(clientCalc, -243)

}
