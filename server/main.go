package main

import (
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

// func (s *server) Add(_ context.Context, request *proto.Request) (*proto.Response, error) {
// 	a, b := request.GetA(), request.GetB()

// 	result := a + b

// 	return &proto.Response{Result: result}, nil
// }

// func (s *server) Multiply(_ context.Context, request *proto.Request) (*proto.Response, error) {
// 	a, b := request.GetA(), request.GetB()

// 	result := a * b

// 	return &proto.Response{Result: result}, nil
// }

// func (s *server) Login(_ context.Context, request *proto.Request) (*proto.Response, error) {
// 	email, password := request.
// }