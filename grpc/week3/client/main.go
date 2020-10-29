package main

import (
	"fmt"

	"github.com/hieuthien95/golang/grpc/week3"
	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("0.0.0.0:19952", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Dial: " + err.Error())
	}
	defer cc.Close()

	client := week3.NewCalculatorServiceClient(cc)

	fmt.Printf("NewCalculatorServiceClient %f", client)
}
