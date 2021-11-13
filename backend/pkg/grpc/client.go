package grpc

import (
	pb "github.com/htchan/UserService/backend/internal/grpc"
	grpc "google.golang.org/grpc"
)

func NewClient(address string) pb.UserServiceClient {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	return pb.NewUserServiceClient(conn)
}