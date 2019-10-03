package connections

import pb "projects/grpc-connections/internal/grpc/connection"

const MaxServerConnectionsQuantity int = 4

func GetServerConnectionsData() []*pb.Connection {
	connections := []*pb.Connection{
		{Id: 1, Host: "localhost", Port: "7000"},
		{Id: 2, Host: "localhost", Port: "7001"},
		{Id: 3, Host: "localhost", Port: "7002"},
		{Id: 4, Host: "localhost", Port: "7003"},
	}

	return connections
}
