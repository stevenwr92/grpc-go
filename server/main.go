package main

import (
	"context"
	"net"

	"golangtest.com/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.AddServiceAuthServer(srv, &server{})
	proto.AddServiceUsersServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(lis); e != nil {
		panic(err)
	}
}

func (s *server) Login(_ context.Context, request *proto.Request) (*proto.Response, error) {
	
	return request.
	
}


func (s *server) GetAllUser(_ context.Context, request *proto.Request) (*proto.Response, error) {
	
	return request.
	
}


func (s *server) GetAllUser(_ context.Context, request *proto.Request) (*proto.Response, error) {
	
	return request.
	
}
func (s *server) CreateUser(_ context.Context, request *proto.Request) (*proto.Response, error) {
	
	return request.
	
}


func (s *server) UpdateUser(_ context.Context, request *proto.Request) (*proto.Response, error) {
	
	return request.
	
}


func (s *server) DeleteUser(_ context.Context, request *proto.Request) (*proto.Response, error) {
	
	return request.
	
}


