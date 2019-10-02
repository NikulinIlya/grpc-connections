package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "projects/grpc-connections/internal/gRPC/connection"
)

const (
	portOne   = "7000"
	portTwo   = "7001"
	portThree = "7002"
)

func main() {
	serverAddress := "localhost:" + portThree

	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())

	if e != nil {
		panic(e)
	}
	defer conn.Close()

	client := pb.NewConnectionServiceClient(conn)

	for i := range [10]int{} {
		clickModel := pb.Click{
			TdsId: uint32(i),
			Id:    uint64(i),
			Ip:    uint32(i),
			Query: []byte(string(i)),
		}

		if responseMessage, e := client.SendClick(context.Background(), &clickModel); e != nil {
			panic(fmt.Sprintf("Was not able to send Click %v", e))
		} else {
			fmt.Println("Click Sent...")
			fmt.Println(responseMessage)
			fmt.Println("=============================")
		}
	}
}
