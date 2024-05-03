package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	pb "your-package-path/user"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)

	// Call AddUser
	addUserResponse, err := c.AddUser(context.Background(), &pb.User{Id: 1, Name: "John", Email: "john@example.com"})
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	log.Printf("User added: %v", addUserResponse)

	// Call GetUser
	getUserResponse, err := c.GetUser(context.Background(), &pb.UserID{Id: 1})
	if err != nil {
		log.Fatalf("could not get user: %v", err)
	}
	log.Printf("User retrieved: %v", getUserResponse)

	// Call ListUsers
	listUsersResponse, err := c.ListUsers(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	log.Printf("Users list: %v", listUsersResponse)
}
