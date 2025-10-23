package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/duonghuu/go_starter_01/userpb"
	"google.golang.org/grpc"
)

type server struct {
	userpb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *userpb.UserRequest) (*userpb.UserResponse, error) {
	return &userpb.UserResponse{Id: req.Id, Name: "Alice"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	userpb.RegisterUserServiceServer(s, &server{})
	fmt.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
