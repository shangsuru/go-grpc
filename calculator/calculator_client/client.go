package main

import (
	"context"
	"fmt"
	"log"

	"github.com/shangsuru/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)
	fmt.Println("Created client")

	res, err2 := c.Calculator(context.Background(), &calculatorpb.Summands{A: 4, B: 5})
	if err2 != nil {
		log.Fatalf("Error when getting result from server: %v", err2)
	}

	fmt.Println("The result is: ", res.GetResult())
}
