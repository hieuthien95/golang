package main

import (
	"context"
	"fmt"
	"io"

	"github.com/hieuthien95/golang/grpc/week5"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:19952", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial: " + err.Error())
	}
	defer cc.Close()

	client := week5.NewSSServiceClient(cc)

	// call SS
	err = sendSS(client)
	fmt.Println(err)
}

func sendSS(client week5.SSServiceClient) error {
	fmt.Println("client streaming")

	stream, err := client.SSFunc(context.Background(), &week5.SSRequest{Text: "abcdef"})
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("end of stream")
			break
		}

		fmt.Println("Output: " + resp.GetOutput())
	}

	return err
}
