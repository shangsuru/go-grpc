package main

import (
	"fmt"
	"io"
	"log"
	"net"

	"github.com/shangsuru/go-grpc/average/averagepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) ComputeAverage(stream averagepb.AverageService_ComputeAverageServer) error {
	fmt.Println("Got RCP request to compute an average")
	var sum, counter int32 = 0, 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("Sending back the average")
			return stream.SendAndClose(&averagepb.Average{Result: sum / counter})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		fmt.Println("Receiving ", req.GetN())
		sum += req.GetN()
		counter++
	}
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	averagepb.RegisterAverageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
