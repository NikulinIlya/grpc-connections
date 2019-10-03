package impl

import (
	"context"
	"log"
	"projects/grpc-connections/internal/grpc/connection"
)

//ConnectionServiceGrpcImpl is a implementation of ConnectionService Grpc Service.
type ConnectionServiceGrpcImpl struct{}

//NewConnectionServiceGrpcImpl returns the pointer to the implementation.
func NewConnectionServiceGrpcImpl() *ConnectionServiceGrpcImpl {
	return &ConnectionServiceGrpcImpl{}
}

func (serviceImpl *ConnectionServiceGrpcImpl) SendClick(ctx context.Context, in *connection.ClickData) (*connection.ClickResponse, error) {
	log.Println("Received data:", in.Id, in.Connection.Host, in.Connection.Port, "Click Id:", in.Click.Id)

	return &connection.ClickResponse{
		Click:      in.Click,
		Connection: in.Connection,
		Error:      nil,
	}, nil
}
