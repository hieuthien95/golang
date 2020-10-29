package main

import (
	"context"
	"fmt"
	"net"

	"github.com/hieuthien95/golang/grpc/week4"
	"google.golang.org/grpc"
)

// impl CalculatorServiceServer
type server struct{}

// Sum
func (*server) Sum(ctx context.Context, req *week4.SumRequest) (*week4.SumResponse, error) {
	fmt.Println("server sum...")

	resp := &week4.SumResponse{
		Result: req.GetNum1() + req.GetNum2(),
	}

	return resp, nil
}

func main() {
	// new listener
	lis, err := net.Listen("tcp", "0.0.0.0:19952")
	if err != nil {
		fmt.Println("Listen: " + err.Error())
		return
	}

	s := grpc.NewServer()

	week4.RegisterCalculatorServiceServer(s, &server{})

	// run
	fmt.Println("Runing...")

	err = s.Serve(lis)
	if err != nil {
		fmt.Println("Serve: " + err.Error())
	}
}
