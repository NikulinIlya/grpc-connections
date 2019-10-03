package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	c "projects/grpc-connections/cmd/grpc/connections"
	pb "projects/grpc-connections/internal/grpc/connection"
)

const CurrentConnectionNumber = 0

func MakeClickModel(connection *pb.Connection) pb.ClickData {
	clickModel := pb.ClickData{
		Id: 1,
		Click: &pb.Click{
			TdsId: 100,
			Id:    50,
			Ip:    10000,
			Query: []byte{
				10,
				20,
			},
		},
		Connection: connection,
	}

	return clickModel
}

func main() {
	var connections [c.MaxServerConnectionsQuantity]*grpc.ClientConn

	serverConnectionsData := c.GetServerConnectionsData()

	for i := range [c.MaxServerConnectionsQuantity]int{} {
		host := serverConnectionsData[i].GetHost()
		port := serverConnectionsData[i].GetPort()

		serverAddress := host + ":" + port

		fmt.Println("Added server address:", host, port)

		conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())
		if e != nil {
			panic(e)
		}

		connections[i] = conn
		defer connections[i].Close()
	}

	fmt.Println("=============================")

	client := pb.NewConnectionServiceClient(connections[CurrentConnectionNumber])

	connection := &pb.Connection{
		Id:   serverConnectionsData[CurrentConnectionNumber].GetId(),
		Host: serverConnectionsData[CurrentConnectionNumber].GetHost(),
		Port: serverConnectionsData[CurrentConnectionNumber].GetPort(),
	}

	clickModel := MakeClickModel(connection)

	if responseMessage, e := client.SendClick(context.Background(), &clickModel); e != nil {
		panic(fmt.Sprintf("Was not able to send Click %v", e))
	} else {
		fmt.Println("Click Sent...")
		fmt.Println("Response:", responseMessage)
		fmt.Println("=============================")
	}
}
