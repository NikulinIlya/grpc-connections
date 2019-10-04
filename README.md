## GRPC-connections

A microservice in Golang which communicates over gRPC. Makes transferring special data using Unary RPC and Client 
streaming RPC, some time measurements for it in client.

### Commands
- protoc -I internal/proto-files/connection/ internal/proto-files/connection/click.proto --go_out=plugins=grpc:internal
  
- protoc -I internal/proto-files/connection/ internal/proto-files/connection/connection-service.proto --go_out=plugins=grpc:internal
  
- go run .\cmd\gRPC\server\main.go
  
- go run .\cmd\gRPC\client\main.go