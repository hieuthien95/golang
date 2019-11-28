package main

import (
	"context"
	"fmt"

	pb "demo-proto/proto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

func main() {
	testFind()
}

func testCreate() {
	// 1. Connect to server at TCP port
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	// 2. New client
	client := pb.NewNoteServiceClient(conn)
	// 3. Call Create
	req := pb.NoteReq{
		Title:     "Todo 123",
		Completed: true,
	}
	res, _ := client.Create(context.TODO(), &req)
	// 4. In ket qua

	fmt.Println("Response:", res)
	fmt.Println("Response.Completed:", res.Completed)
}

func testFind() {
	// 1. Connect to server at TCP port
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	// 2. New client
	client := pb.NewNoteServiceClient(conn)
	// 3. Call Create
	req := pb.NoteFindReq{
		Id: 12345,
	}
	res, _ := client.Find(context.TODO(), &req)
	// 4. In ket qua

	fmt.Println("Response:", res)
	fmt.Println("Response.Completed:", res.Completed)
}

func testDelete() {
	// 1. Connect den server
	conn, _ := grpc.Dial(address, grpc.WithInsecure())
	client := pb.NewNoteServiceClient(conn)
	req := pb.NoteDelReq{
		Id: 124,
	}
	res, _ := client.Delete(context.TODO(), &req)

	if res.Success == false {
		fmt.Println("Can not delete")
	} else {
		fmt.Println("Can delete")
	}
}
