package main

import (
	"context"
	"google.golang.org/grpc"
	"todoClient/proto"
)

func main() {
	con, err := grpc.Dial("localhost:4040", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	client := proto.NewAddServiceClient(con)

	req := proto.Request{A:2, B:4}

	ctx := context.Background()

	response, err := client.Multiply(ctx, &req)

	if err != nil {
		panic(err)
	}

	println(response.Result)
}
