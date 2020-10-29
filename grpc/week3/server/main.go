package main

import (
	"fmt"
	"net"

	"github.com/hieuthien95/golang/grpc/week3"
	"google.golang.org/grpc"
)

// impl CalculatorServiceServer
type server struct{}

func main() {
	// new listener
	lis, err := net.Listen("tcp", "0.0.0.0:19952")
	if err != nil {
		fmt.Println("Listen: " + err.Error())
		return
	}

	s := grpc.NewServer()

	week3.RegisterCalculatorServiceServer(s, &server{})

	// run
	fmt.Println("Runing...")

	err = s.Serve(lis)
	if err != nil {
		fmt.Println("Serve: " + err.Error())
	}
}
