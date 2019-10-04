package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"math/rand"
	c "projects/grpc-connections/cmd/grpc/connections"
	pb "projects/grpc-connections/internal/grpc/connection"
	"time"
)

const (
	CurrentConnectionNumber = 0
	ClickArraySize          = 1000
)

func MakeClickModel(connection *pb.Connection, Id int64) pb.ClickData {
	clickModel := pb.ClickData{
		Id: Id,
		Click: &pb.Click{
			TdsId: rand.Uint32(),
			Id:    rand.Uint64(),
			Ip:    rand.Uint32(),
			Query: []byte{
				byte(rand.Uint32()),
				byte(rand.Uint32()),
				byte(rand.Uint32()),
			},
		},
		Connection: connection,
	}

	return clickModel
}

func runClickSend(client pb.ConnectionServiceClient, clickModelArr [ClickArraySize]pb.ClickData) {
	fmt.Println("=============================")
	fmt.Println("ClickSend started")
	t0 := time.Now()
	for range [1000]int{} {
		//t1 := time.Now()

		for i := range [1000]int{} {

			//responseMessage
			if _, e := client.SendClick(context.Background(), &clickModelArr[i]); e != nil {
				panic(fmt.Sprintf("Was not able to send Click %v", e))
			} else {
				/*fmt.Println("Click Sent...")
				fmt.Println("Response:", responseMessage)
				fmt.Println("=============================")*/
			}
		}
		//fmt.Printf("Elapsed time on %d: %v \n", j, time.Since(t1))
	}

	fmt.Printf("Sum Elapsed time: %v \n", time.Since(t0))
	fmt.Printf("Average 1000 Send: %v \n", time.Since(t0)/1000)
	fmt.Printf("Average one Send: %v \n", time.Since(t0)/1000000)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

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

	var clickModelArr [ClickArraySize]pb.ClickData

	for i := range [ClickArraySize]int{} {
		clickModelArr[i] = MakeClickModel(connection, int64(i))
	}

	fmt.Println("clickModelArr generated")

	runClickSend(client, clickModelArr)
}
