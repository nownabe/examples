package main

import (
	"google.golang.org/grpc"

	greetv1 "github.com/nownabe/examples/2023-11-18-connect-go/gen/greet/v1"
)

func NewGreetServiceClient() greetv1.GreetServiceClient {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	return greetv1.NewGreetServiceClient(conn)
}
