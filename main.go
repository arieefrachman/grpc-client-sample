package main

import (
	"context"
	"grpc-client-sample/proto"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:7001", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(con)

	req := proto.Request{A: 2, B: 4}

	ctx := context.Background()

	response, err := client.Multiply(ctx, &req)

	if err != nil {
		panic(err)
	}

	println(response.Result)
}
