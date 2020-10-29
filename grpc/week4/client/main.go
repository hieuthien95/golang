package main

import (
	"context"
	"fmt"

	"github.com/hieuthien95/golang/grpc/week4"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:19952", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial: " + err.Error())
	}
	defer cc.Close()

	client := week4.NewCalculatorServiceClient(cc)

	// call sum
	sum, err := callSum(client)
	fmt.Println(sum, err)
}

func callSum(client week4.CalculatorServiceClient) (int32, error) {
	fmt.Println("client calling sum")

	resp, err := client.Sum(context.Background(), &week4.SumRequest{Num1: 10, Num2: 20})
	if err != nil {
		return 0, err
	}

	return resp.GetResult(), err
}
