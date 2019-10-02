package impl

import (
	"context"
	"log"
	"projects/grpc-connections/internal/gRPC/connection"
	"strconv"
)

//ConnectionServiceGrpcImpl is a implementation of ConnectionService Grpc Service.
type ConnectionServiceGrpcImpl struct{}

//NewConnectionServiceGrpcImpl returns the pointer to the implementation.
func NewConnectionServiceGrpcImpl() *ConnectionServiceGrpcImpl {
	return &ConnectionServiceGrpcImpl{}
}

//Add function implementation of gRPC Service.
func (serviceImpl *ConnectionServiceGrpcImpl) SetConnection(ctx context.Context, in *connection.Connection) (*connection.AddConnectionResponse, error) {
	log.Println("Received request for setting connection with id " + strconv.FormatInt(in.Id, 10))

	log.Println("Connected")

	return &connection.AddConnectionResponse{
		AddedConnection: in,
		Error:           nil,
	}, nil
}

func (serviceImpl *ConnectionServiceGrpcImpl) SendClick(ctx context.Context, in *connection.Click) (*connection.ClickResponse, error) {
	log.Println("Received data: ", in.Id)

	return &connection.ClickResponse{
		Click: in,
		Error: nil,
	}, nil
}
