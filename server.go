package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "your-package-path/user"
)

type server struct{}

func (s *server) AddUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	// Add user logic here
	return user, nil
}

func (s *server) GetUser(ctx context.Context, userID *pb.UserID) (*pb.User, error) {
	// Get user l
