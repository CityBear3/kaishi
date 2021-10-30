package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	proto "kaishi"
	"log"
	"net"
)

type Server struct {
	proto.UnimplementedGreetServiceServer
}

func (*Server) Greet(ctx context.Context, request *proto.GreetRequest) (*proto.GreetResponse, error) {
	name := request.GetName()
	message := fmt.Sprintf("Hello, %s", name)

	return &proto.GreetResponse{
		Greet: message,
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", "localhost:58080")
	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)

	proto.RegisterGreetServiceServer(server, &Server{})

	reflection.Register(server)
	if e := server.Serve(listen); e != nil {
		log.Fatalln(e)
	}
}
