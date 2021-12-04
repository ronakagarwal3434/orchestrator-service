package main

import (
	"context"
	"errors"
	"net"

	"orchestrator-servic/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
}

func (s *server) GetUserByName(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	return nil, errors.New("not implemented yet. <yourname> will implement me")
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	proto.RegisterServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
