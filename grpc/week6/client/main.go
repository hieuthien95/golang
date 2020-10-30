package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hieuthien95/golang/grpc/week6"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:19952", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial: " + err.Error())
	}
	defer cc.Close()

	client := week6.NewCSServiceClient(cc)

	// call CS
	err = sendCS(client)
	fmt.Println(err)
}

func sendCS(client week6.CSServiceClient) error {
	fmt.Println("client streaming")

	// open stream
	stream, err := client.CSFunc(context.Background())
	if err != nil {
		return err
	}

	// send message
	messages := []string{"1", "2", "3", "4"}
	for _, s := range messages {
		err := stream.Send(&week6.CSRequest{
			Text: s,
		})

		if err != nil {
			return err
		}

		time.Sleep(time.Second)
	}

	// receive
	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	fmt.Println("CloseAndRecv: " + resp.GetOutput())

	return err
}
