protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
protoc --go_out=./hello --go_opt=paths=source_relative --go-grpc_out=./hello --go-grpc_opt=paths=source_relative hello.proto
protoc --go_out=. --go-grpc_out=. hello.proto

to start the gRPC server:
go run server/server.go

in a separate terminal window, to execute the gRPC client request:
go run server/server.go

PostMan:
{
    "name": "Dipu"
}

Follow this as a ref: https://pascalallen.medium.com/how-to-build-a-grpc-server-in-go-943f337c4e05
