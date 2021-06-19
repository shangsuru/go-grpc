package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/shangsuru/go-grpc/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Calculator(ctx context.Context, summands *calculatorpb.Summands) (*calculatorpb.Sum, error) {
	fmt.Printf("Calculating the sum of %d and %d ...", summands.GetA(), summands.GetB())
	return &calculatorpb.Sum{
		Result: summands.GetA() + summands.GetB(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: $%v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
