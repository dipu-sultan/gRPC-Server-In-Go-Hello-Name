protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/*.proto
protoc --go_out=./hello --go_opt=paths=source_relative --go-grpc_out=./hello --go-grpc_opt=paths=source_relative hello.proto
protoc --go_out=. --go-grpc_out=. hello.proto

PostMan:
{
    "name": "Dipu"
}
