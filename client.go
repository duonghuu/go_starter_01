package main

import (
	"context"
	"fmt"
	"log"

	"github.com/duonghuu/go_starter_01/userpb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	c := userpb.NewUserServiceClient(conn)
	res, err := c.GetUser(context.Background(), &userpb.UserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("Response: %v\n", res)
}
