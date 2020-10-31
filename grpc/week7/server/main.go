package main

import (
	"fmt"
	"io"
	"net"

	"github.com/hieuthien95/golang/grpc/week7"
	"google.golang.org/grpc"
)

// impl CalculatorServiceServer
type server struct{}

func (*server) MinMax2WayStreaming(stream week7.CalculatorService_MinMax2WayStreamingServer) error {
	var min int32 = 9999999
	var max int32 = 0
	for {
		// GetInput
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println(err)
			break
		}
		if err != nil {
			fmt.Println(err)
			return err
		}

		input := req.GetInput()
		fmt.Println(input)

		// check
		if input < min {
			min = input
		}
		if input > max {
			max = input
		}

		// Send
		err = stream.Send(&week7.MinMaxResponse{
			Result: fmt.Sprintf("Min: %v", min),
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
		err = stream.Send(&week7.MinMaxResponse{
			Result: fmt.Sprintf("Max: %v", max),
		})
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	return nil
}

func main() {
	// new listener
	lis, err := net.Listen("tcp", "0.0.0.0:19952")
	if err != nil {
		fmt.Println("Listen: " + err.Error())
		return
	}

	s := grpc.NewServer()

	week7.RegisterCalculatorServiceServer(s, &server{})

	// run
	fmt.Println("Runing...")

	err = s.Serve(lis)
	if err != nil {
		fmt.Println("Serve: " + err.Error())
	}
}
