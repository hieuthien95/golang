package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/hieuthien95/golang/grpc/week5"
	"google.golang.org/grpc"
)

// impl CalculatorServiceServer
type server struct{}

// Sum
func (*server) SSFunc(req *week5.SSRequest, stream week5.SSService_SSFuncServer) error {
	str := req.Text
	arrStr := strings.Split(str, "")

	for _, s := range arrStr {
		err := stream.Send(&week5.SSResponse{
			Output: s,
		})

		if err != nil {
			fmt.Println("Server Send: " + err.Error())
		}

		time.Sleep(time.Second)
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

	week5.RegisterSSServiceServer(s, &server{})

	// run
	fmt.Println("Runing...")

	err = s.Serve(lis)
	if err != nil {
		fmt.Println("Serve: " + err.Error())
	}
}
