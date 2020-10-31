package main

import (
	"context"
	"fmt"
	"io"
	"sync"
	"time"

	"github.com/hieuthien95/golang/grpc/week7"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:19952", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial: " + err.Error())
	}
	defer cc.Close()

	client := week7.NewCalculatorServiceClient(cc)

	// do action
	err = sendInput(client)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func sendInput(client week7.CalculatorServiceClient) error {
	stream, err := client.MinMax2WayStreaming(context.Background())
	if err != nil {
		return err
	}

	arrInt := []int32{4, 5, 6, 7, 1, 2, 3}

	wg := new(sync.WaitGroup)
	wg.Add(2)

	// send
	go func() {
		// CloseSend
		defer stream.CloseSend()
		defer wg.Done()

		// Send
		for _, num := range arrInt {
			err := stream.Send(&week7.MinMaxRequest{
				Input: num,
			})
			if err != nil {
				if err != nil {
					fmt.Println(err.Error())
				}
				break
			}

			time.Sleep(time.Second)
		}
	}()

	// receive
	go func() {
		defer wg.Done()

		for {
			// Recv
			resp, err := stream.Recv()
			if err == io.EOF {
				fmt.Println(err)
				return
			}
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			result := resp.GetResult()
			fmt.Println(result)
		}
	}()

	wg.Wait()

	return nil
}
