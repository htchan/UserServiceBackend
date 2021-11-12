package grpc

import (
	"context"
	goGrpc"google.golang.org/grpc"
	"github.com/htchan/UserService/internal/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"github.com/htchan/UserService/pkg/users"
	"github.com/htchan/UserService/pkg/services"
	"github.com/htchan/UserService/pkg/tokens"
	"github.com/htchan/UserService/pkg/permissions"
	"log"
	"net"
	"fmt"
)

type Server struct {
	grpc.UnimplementedUserServiceServer
}


// Authenticate(ctx context.Context, in *grpc.TokenWithPermission) (*grpc.Result, error)
// Authorize(ctx context.Context, in *grpc.AuthorizeParams) (*grpc.Result, error)
// RegisterPermission(ctx context.Context, in *grpc.TokenWithPermission) (*grpc.Result, error)
// UnregisterPermission(ctx context.Context, in *grpc.TokenWithPermission) (*grpc.Result, error)

func (server *Server) Signup(ctx context.Context, in *grpc.LoginParams) (*grpc.AuthToken, error) {
	user, err := users.Signup(*in.Username, *in.Password)
	if err != nil {
		return nil, err
	}
	userToken, err := tokens.LoadUserToken(*user, 60*60*24)
	if err != nil {
		return nil, err
	}
	token := new(grpc.AuthToken)
	token.Token = &userToken.Token
	return token, nil
}

func (server *Server) DropOut(ctx context.Context, in *grpc.AuthToken) (*grpc.Result, error) {
	user, err := tokens.FindUserByTokenStr(*in.Token)
	if err != nil {
		return nil, err
	}
	// remove user's permission
	userPermissions, err := permissions.FindUserPermissionsByUser(*user)
	if err != nil {
		return nil, err
	}
	for _, permission := range userPermissions {
		permissions.RevokePermission(permission)
	}
	// remove user's auth token
	tokens.DeleteUserTokens(*user)
	// remove user in db
	err = users.Dropout(*user)
	if err != nil {
		return nil, err
	}
	s := "success"
	return &grpc.Result{Result: &s}, nil
}

func (server *Server) Login(ctx context.Context, in *grpc.LoginParams) (*grpc.AuthToken, error) {
	// find / generate token for user
	user, err := users.Login(*in.Username, *in.Password)
	if err != nil {
		return nil, err
	}
	userToken, err := tokens.LoadUserToken(*user, 24 * 60)
	if err != nil {
		return nil, err
	}
	token := new(grpc.AuthToken)
	token.Token = &userToken.Token
	return token, nil
}

func (server *Server) Logout(ctx context.Context, in *grpc.AuthToken) (*emptypb.Empty, error) {
	token, err := tokens.FindUserTokenByTokenStr(*in.Token)
	if err != nil {
		return nil, err
	}
	err = token.Expire()
	if err != nil {
		return nil, err
	}
	return new(emptypb.Empty), nil
}

func (server *Server) RegisterService(ctx context.Context, in *grpc.ServiceName) (*grpc.AuthToken, error) {
	service, err := services.RegisterService(*in.Name)
	if err != nil {
		return nil, err
	}
	serviceToken, err := tokens.LoadServiceToken(*service)
	if err != nil {
		return nil, err
	}
	token := new(grpc.AuthToken)
	token.Token = &serviceToken.Token
	return token, nil
}

func (server *Server) UnregisterService(ctx context.Context, in *grpc.AuthToken) (*grpc.Result, error) {
	service, err := tokens.FindServiceByTokenStr(*in.Token)
	if err != nil {
		return nil, err
	}
	err = services.UnregisterService(*service)
	if err != nil {
		return nil, err
	}
	s := "success"
	return &grpc.Result{Result: &s}, nil
}

func (server *Server) RegisterPermission(ctx context.Context, in *grpc.TokenWithPermission) (*grpc.Result, error) {
	service, err := tokens.FindServiceByTokenStr(*in.Token)
	if err != nil {
		return nil, err
	}
	_, err = permissions.RegisterPermission(*service, *in.Permission)
	if err != nil {
		return nil, err
	}
	s := "success"
	return &grpc.Result{Result: &s}, nil
}

func (server *Server) UnregisterPermission(ctx context.Context, in *grpc.TokenWithPermission) (*grpc.Result, error) {
	service, err := tokens.FindServiceByTokenStr(*in.Token)
	if err != nil {
		return nil, err
	}
	servicePermission, err := permissions.FindServicePermissionByPermission(*service, *in.Permission)
	if err != nil {
		return nil, err
	}
	err = permissions.UnregisterPermission(*service, *servicePermission)
	if err != nil {
		return nil, err
	}
	s := "success"
	return &grpc.Result{Result: &s}, nil
}


func (server *Server) Authenticate(ctx context.Context, in *grpc.TokenWithPermission) (*grpc.Result, error) {
	// check user has permission
	user, err := tokens.FindUserByTokenStr(*in.Token)
	if err != nil {
		return nil, err
	}
	_, err = permissions.FindUserPermissionByPermission(*user, *in.Permission)
	if err != nil {
		return nil, err
	}
	s := "success"
	return &grpc.Result{Result: &s}, nil
}

func (server *Server) Authorize(ctx context.Context, in *grpc.AuthorizeParams) (*grpc.Result, error) {
	// give permission to user
	service, err := tokens.FindServiceByTokenStr(*in.Token)
	if err != nil {
		return nil, err
	}
	servicePermission, err := permissions.FindServicePermissionByPermission(*service, *in.Permission)
	if err != nil {
		return nil, err
	}
	user, err := users.FindUserByName(*in.Username)
	if err != nil {
		return nil, err
	}
	permissions.GrantPermission(*user, *servicePermission)
	s := "success"
	return &grpc.Result{Result: &s}, nil
}

func StartServer(port int) {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%v", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := goGrpc.NewServer()
	grpc.RegisterUserServiceServer(s, &Server{})
	log.Printf("server listening at %v", listen.Addr())
	log.Println("grpc start")
	log.Fatal(s.Serve(listen))
}