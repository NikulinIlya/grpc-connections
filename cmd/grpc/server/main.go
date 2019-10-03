package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	c "projects/grpc-connections/cmd/grpc/connections"
	pb "projects/grpc-connections/internal/grpc/connection"
	"projects/grpc-connections/internal/grpc/impl"
)

func runServe(port string) {
	netListenerOne := getNetListener(port)

	gRPCServer := grpc.NewServer()

	connectionServiceImpl := impl.NewConnectionServiceGrpcImpl()
	pb.RegisterConnectionServiceServer(gRPCServer, connectionServiceImpl)

	// start the server
	log.Println("Start Server, port:", port)
	if err := gRPCServer.Serve(netListenerOne); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func getNetListener(port string) net.Listener {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}

func main() {
	connections := c.GetServerConnectionsData()

	for i := range [c.MaxServerConnectionsQuantity - 1]int{} {
		go runServe(connections[i].GetPort())
	}

	runServe(connections[c.MaxServerConnectionsQuantity-1].GetPort())
}
