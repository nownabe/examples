package main

import (
	"fmt"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func echo(ctx context.Context, in *Request) (*Response, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	fmt.Println(", Metadata:", md)
	return &Response{Message: in.Message}, nil
}

type server struct{}

func (server) Echo1(ctx context.Context, in *Request) (*Response, error) {
	fmt.Print("Echo1 Received: ", in.Message)
	return echo(ctx, in)
}

func (server) Echo2(ctx context.Context, in *Request) (*Response, error) {
	fmt.Print("Echo2 Received: ", in.Message)
	return echo(ctx, in)
}

func (server) Echo3(ctx context.Context, in *Request) (*Response, error) {
	fmt.Print("Echo3 Received: ", in.Message)
	return echo(ctx, in)
}

func (server) Echo4(ctx context.Context, in *Request) (*Response, error) {
	fmt.Print("Echo4 Received: ", in.Message)
	return echo(ctx, in)
}

func main() {
	s := grpc.NewServer()
	RegisterEchoServiceServer(s, server{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
