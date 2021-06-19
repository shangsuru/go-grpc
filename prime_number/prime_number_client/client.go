package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/shangsuru/go-grpc/prime_number/primepb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := primepb.NewPrimeNumberDecompositionClient(conn)
	fmt.Println("Created client")

	fmt.Println("Make Decompose RCP call")
	resStream, err2 := c.Decompose(context.Background(), &primepb.Number{Number: 156})
	if err2 != nil {
		log.Fatalf("Error when getting result from server: %v", err2)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("Reached end of stream")
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		fmt.Println("Received prime factor: ", msg.GetPrime())
	}
}
