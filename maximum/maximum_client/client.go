package main

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"time"

	"github.com/shangsuru/go-grpc/maximum/maximumpb"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()

	c := maximumpb.NewMaximumServiceClient(conn)

	requests := []*maximumpb.Number{{Number: 3}, {Number: 2}, {Number: 11}, {Number: 13}}

	stream, _ := c.GetMaximum(context.Background())

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			fmt.Println("The current maximum is ", res.GetMaximum())
		}
	}()

	<-waitc
}
