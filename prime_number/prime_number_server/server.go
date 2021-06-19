package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/shangsuru/go-grpc/prime_number/primepb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Decompose(number *primepb.Number, stream primepb.PrimeNumberDecomposition_DecomposeServer) error {
	fmt.Println("Received RPC call for Decompose of ", number.GetNumber())

	N := number.GetNumber()
	var k int32 = 2
	for N > 1 {
		if N%k == 0 {
			fmt.Println("Factor found: ", k)
			N = N / k

			stream.Send(&primepb.PrimeFactor{Prime: k})
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		} else {
			k++
		}
	}

	fmt.Println("Returning")
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: $%v", err)
	}

	s := grpc.NewServer()
	primepb.RegisterPrimeNumberDecompositionServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
