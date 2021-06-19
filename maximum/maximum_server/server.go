package main

import (
	"fmt"
	"io"
	"math/rand"
	"net"
	"time"

	"github.com/shangsuru/go-grpc/maximum/maximumpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) GetMaximum(stream maximumpb.MaximumService_GetMaximumServer) error {
	fmt.Println("Rerceiving RPC call")

	var max int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}

		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		n := req.GetNumber()
		fmt.Println("Receiving ", n)

		if n > max {
			max = n
		}

		fmt.Println("New maximum is ", max)
		stream.Send(&maximumpb.Maximum{Maximum: max})
	}

	fmt.Println("End RPC call")
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", "0.0.0.0:50051")
	s := grpc.NewServer()
	maximumpb.RegisterMaximumServiceServer(s, &server{})
	s.Serve(lis)
}
