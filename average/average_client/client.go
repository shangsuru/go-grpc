package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/shangsuru/go-grpc/average/averagepb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := averagepb.NewAverageServiceClient(conn)
	fmt.Println("Created client connection")

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverage: %v", err)
	}

	requests := []*averagepb.Number{
		{N: 4}, {N: 45}, {N: 32}, {N: 9}, {N: 11}, {N: 65},
	}

	for _, req := range requests {
		fmt.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error receiving result: %v", err)
	}

	fmt.Println("The average is: ", res.GetResult())
}
