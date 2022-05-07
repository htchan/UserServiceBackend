package grpc

import (
	"context"
	goGrpc "google.golang.org/grpc"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/internal/grpc"
	"github.com/htchan/UserService/backend/pkg/user"
	"github.com/htchan/UserService/backend/pkg/service"
	"github.com/htchan/UserService/backend/pkg/token"
	"github.com/htchan/UserService/backend/pkg/permissions"
	"log"
	"net"
	"errors"
)

type Server struct {
	grpc.UnimplementedUserServiceServer
}

func recoverError() {
	recover()
}

func (server *Server) Signup(ctx context.Context, in *grpc.SignupParams) (authToken *grpc.AuthToken, err error) {
	userToken, err := token.UserSignup(*in.Username, *in,Password)
	authToken = new(grpc.AuthToken)
	authToken.Token = &userToken.Token
	return
}

func (server *Server) Dropout(ctx context.Context, in *grpc.AuthToken) (result *grpc.Result, err error) {
	err = token.UserDropout(*in.Token)
	msg := "failed"
	if err != nil {
		msg = "success"
	}
	result = &grpc.Result{Result: &msg}
	return
}

func (server *Server) Login(ctx context.Context, in *grpc.LoginParams) (tokenWithUrl *grpc.TokenWithUrl, err error) {
	userToken, url, err := token.UserNameLogin(
		*in.Username,
		*in.Password,
		*in.Service,
	)

	tokenWithUrl = new(grpc.TokenWithUrl)
	tokenWithUrl.Token = &userToken.Token
	tokenWithUrl.Url = &url
	return
}

func (server *Server) Logout(ctx context.Context, in *grpc.AuthToken) (result *grpc.Result, err error) {
	err = token.UserLogout(*in.Token)
	msg := "failed"
	if err != nil {
		msg = "success"
	}
	result = &grpc.Result{Result: &msg}
	return
}

func (server *Server) RegisterService(ctx context.Context, in *grpc.ServiceName) (authToken *grpc.AuthToken, err error) {
	serviceToken, err := token.ServiceRegister(*in.Name, *in.Url)
	authToken = new(grpc.AuthToken)
	authToken.Token = &serviceToken.Token
	return
}

func (server *Server) UnregisterService(ctx context.Context, in *grpc.AuthToken) (result *grpc.Result, err error) {
	err = token.ServiceUnregister(*in.Token)
	msg := "failed"
	if err != nil {
		msg = "success"
	}
	result = &grpc.Result{Result: &s}
	return
}

func (server *Server) RegisterPermission(ctx context.Context, in *grpc.TokenWithPermission) (result *grpc.Result, err error) {
	// defer recoverError()
	// service, err := tokens.FindServiceByTokenStr(*in.Token)
	// utils.CheckError(err)
	// _, err = permissions.RegisterPermission(service, *in.Permission)
	// utils.CheckError(err)
	// s := "success"
	// result = &grpc.Result{Result: &s}
	return
}

func (server *Server) UnregisterPermission(ctx context.Context, in *grpc.TokenWithPermission) (result *grpc.Result, err error) {
	// defer recoverError()
	// service, err := tokens.FindServiceByTokenStr(*in.Token)
	// utils.CheckError(err)
	// servicePermission, err := permissions.FindServicePermissionByPermission(service, *in.Permission)
	// utils.CheckError(err)
	// err = permissions.UnregisterPermission(service, servicePermission)
	// utils.CheckError(err)
	// s := "success"
	// result = &grpc.Result{Result: &s}
	return
}

func (server *Server) Authenticate(ctx context.Context, in *grpc.AuthenticateParams) (result *grpc.Result, err error) {
	// defer recoverError()
	// userToken, err := tokens.FindUserTokenByTokenStr(*in.UserToken)
	// utils.CheckError(err)
	// service, err := tokens.FindServiceByTokenStr(*in.ServiceToken)
	// utils.CheckError(err)
	// if !userToken.BelongsToService(service) {
	// 	err := errors.New("invalid_token")
	// 	utils.CheckError(err)
	// }
	// user, err := tokens.FindUserByTokenStr(*in.UserToken)
	// utils.CheckError(err)
	// if *in.Permission != "" {
	// 	_, err = permissions.FindUserPermissionByPermission(user, *in.Permission)
	// 	utils.CheckError(err)
	// }
	// result = &grpc.Result{Result: &user.UUID}
	tkn, err := token.GetUserToken(*in.Token)
	if err != nil {
		err = errors.New("unauthorized user")
	}
	u, err := tkn.User()
	if err != nil {
		err = errors.New("unauthorized user")
	}
	result = &grpc.Result{Result: &u.UUID}
	return
}

func (server *Server) Authorize(ctx context.Context, in *grpc.AuthorizeParams) (result *grpc.Result, err error) {
	// defer recoverError()
	// // give permission to user
	// service, err := tokens.FindServiceByTokenStr(*in.Token)
	// utils.CheckError(err)
	// servicePermission, err := permissions.FindServicePermissionByPermission(service, *in.Permission)
	// utils.CheckError(err)
	// user, err := users.FindUserByUUID(*in.UserUUID)
	// utils.CheckError(err)
	// err = permissions.GrantPermission(user, servicePermission)
	// utils.CheckError(err)
	// s := "success"
	// result = &grpc.Result{Result: &s}
	return
}

func StartServer(addr string) {
	userService := services.UserService()
	tokens.LoadServiceToken(userService)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := goGrpc.NewServer()
	grpc.RegisterUserServiceServer(s, &Server{})
	log.Printf("server listening at %v", listen.Addr())
	log.Println("grpc start")
	log.Fatal(s.Serve(listen))
}