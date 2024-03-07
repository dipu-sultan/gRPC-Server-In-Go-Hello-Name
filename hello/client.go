package main

import (
	"context"
	"fmt"
	pb "hello/hello"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	address     = "localhost:8080"
	defaultName = "Brooooo"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	client := pb.NewServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Hello(ctx, &pb.Request{Name: name})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res.Res)
}
