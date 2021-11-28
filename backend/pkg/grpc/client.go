package grpc

import (
	pb "github.com/htchan/UserService/backend/internal/grpc"
	grpc "google.golang.org/grpc"
)

type Client interface {
	pb.UserServiceClient
}

func NewClient(address string) Client {
	conn, _ := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	return pb.NewUserServiceClient(conn)
}

func NewAuthToken(token string) *pb.AuthToken {
	return &pb.AuthToken{
		Token: &token,
	}
}

func NewServiceName(name, url string) *pb.ServiceName {
	return &pb.ServiceName{
		Name: &name,
		Url: &url,
	}
}
 
func NewAuthorizeParams(token, userUUID, permission string) *pb.AuthorizeParams {
	return &pb.AuthorizeParams{
		Token: &token,
		UserUUID: &userUUID,
		Permission: &permission,
	}
}

func NewSignupParams(username, password string) *pb.SignupParams {
	return &pb.SignupParams{
		Username: &username,
		Password: &password,
	}
}

func NewLoginParams(username, password, token string) *pb.LoginParams {
	return &pb.LoginParams{
		Username: &username,
		Password: &password,
		Token: &token,
	}
}

func NewTokenWithPermission(token, permission string) *pb.TokenWithPermission {
	return &pb.TokenWithPermission{
		Token: &token,
		Permission: &permission,
	}
}