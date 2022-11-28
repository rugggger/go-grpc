package main

import (
	"context"
	"fmt"

	pb "github.com/rugggger/go-grpc/calculator/proto"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Num: n})
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			fmt.Printf("Error (%d): %s\n", e.Code(), e.Message())
		} else {
			fmt.Printf("Non gRPC error  %v\n", err)
		}
	} else {
		fmt.Printf("Result: %d\n", res.Result)
	}
}
