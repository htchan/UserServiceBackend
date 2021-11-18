package grpc

import (
	pb "github.com/htchan/UserService/backend/internal/grpc"
	grpc "google.golang.org/grpc"
)

func NewClient(address string) pb.UserServiceClient {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	return pb.NewUserServiceClient(conn)
}

func NewAuthToken(token string) *pb.AuthToken {
	return &pb.AuthToken{
		Token: &token,
	}
}

func NewServiceName(name string) *pb.ServiceName {
	return &pb.ServiceName{
		Name: &name,
	}
}
 
func NewAuthorizeParams(token, userUUID, permission string) *pb.AuthorizeParams {
	return &pb.AuthorizeParams{
		Token: &token,
		UserUUID: &userUUID,
		Permission: &permission,
	}
}

func NewLoginParams(username, password string) *pb.LoginParams {
	return &pb.LoginParams{
		Username: &username,
		Password: &password,
	}
}

func NewTokenWithPermission(token, permission string) *pb.TokenWithPermission {
	return &pb.TokenWithPermission{
		Token: &token,
		Permission: &permission,
	}
}