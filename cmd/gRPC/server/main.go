package main

import (
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
	pb "projects/grpc-connections/internal/gRPC/connection"
	"projects/grpc-connections/internal/gRPC/impl"
)

const (
	portOne  = 7000
	portTwo  = 7001
	portTree = 7002
)

func runServe(port uint) {
	netListenerOne := getNetListener(port)

	gRPCServer := grpc.NewServer()

	connectionServiceImpl := impl.NewConnectionServiceGrpcImpl()
	pb.RegisterConnectionServiceServer(gRPCServer, connectionServiceImpl)

	// start the server
	log.Printf("Start Server One: %d", port)
	if err := gRPCServer.Serve(netListenerOne); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func main() {
	go runServe(portOne)
	go runServe(portTwo)
	runServe(portTree)
}

func getNetListener(port uint) net.Listener {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		panic(fmt.Sprintf("failed to listen: %v", err))
	}

	return lis
}
