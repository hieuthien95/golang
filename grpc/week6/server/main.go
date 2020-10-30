package main

import (
	"fmt"
	"io"
	"net"

	"github.com/hieuthien95/golang/grpc/week6"
	"google.golang.org/grpc"
)

// impl CalculatorServiceServer
type server struct{}

// CSFunc
func (*server) CSFunc(stream week6.CSService_CSFuncServer) error {
	str := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("end of stream: " + str)

			// SendAndClose
			stream.SendAndClose(&week6.CSResponse{
				Output: str,
			})
			break
		}
		if err != nil {
			fmt.Println("Recv: " + err.Error())
			return err
		}

		fmt.Println("Text: " + req.GetText())
		str += req.GetText()
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

	week6.RegisterCSServiceServer(s, &server{})

	// run
	fmt.Println("Runing...")

	err = s.Serve(lis)
	if err != nil {
		fmt.Println("Serve: " + err.Error())
	}
}
