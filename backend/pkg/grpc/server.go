package grpc

import (
	"context"
	goGrpc"google.golang.org/grpc"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/internal/grpc"
	"github.com/htchan/UserService/backend/pkg/users"
	"github.com/htchan/UserService/backend/pkg/services"
	"github.com/htchan/UserService/backend/pkg/tokens"
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
	defer recoverError()
	user, err := users.Signup(*in.Username, *in.Password)
	utils.CheckError(err)
	userToken, err := tokens.LoadUserToken(user, services.UserService(), 60*24)
	utils.CheckError(err)
	authToken = new(grpc.AuthToken)
	authToken.Token = &userToken.Token
	return
}

func (server *Server) Dropout(ctx context.Context, in *grpc.AuthToken) (result *grpc.Result, err error) {
	defer recoverError()
	user, err := tokens.FindUserByTokenStr(*in.Token)
	utils.CheckError(err)
	// remove user's permission
	userPermissions, err := permissions.FindUserPermissionsByUser(user)
	utils.CheckError(err)
	for _, permission := range userPermissions {
		permissions.RevokePermission(permission)
	}
	// remove all user's auth token
	tokens.DeleteUserTokens(user)
	// remove user in db
	err = users.Dropout(user)
	utils.CheckError(err)
	s := "success"
	result = &grpc.Result{Result: &s}
	return
}

func (server *Server) Login(ctx context.Context, in *grpc.LoginParams) (tokenWithUrl *grpc.TokenWithUrl, err error) {
	defer recoverError()
	// find / generate token for user
	user, err := users.Login(*in.Username, *in.Password)
	utils.CheckError(err)
	service, err := services.FindServiceByName(*in.Service)
	utils.CheckError(err)
	userToken, err := tokens.LoadUserToken(user, service, 24 * 60)
	utils.CheckError(err)
	tokenWithUrl = new(grpc.TokenWithUrl)
	tokenWithUrl.Token = &userToken.Token
	tokenWithUrl.Url = &service.Url
	return
}

func (server *Server) Logout(ctx context.Context, in *grpc.AuthToken) (result *grpc.Result, err error) {
	defer recoverError()
	token, err := tokens.FindUserTokenByTokenStr(*in.Token)
	utils.CheckError(err)
	err = token.Expire()
	utils.CheckError(err)
	s := "success"
	result = &grpc.Result{Result: &s}
	return
}

func (server *Server) RegisterService(ctx context.Context, in *grpc.ServiceName) (authToken *grpc.AuthToken, err error) {
	defer recoverError()
	service, err := services.RegisterService(*in.Name, *in.Url)
	utils.CheckError(err)
	serviceToken, err := tokens.LoadServiceToken(service)
	utils.CheckError(err)
	authToken = new(grpc.AuthToken)
	authToken.Token = &serviceToken.Token
	return
}

func (server *Server) UnregisterService(ctx context.Context, in *grpc.AuthToken) (result *grpc.Result, err error) {
	defer recoverError()
	service, err := tokens.FindServiceByTokenStr(*in.Token)
	utils.CheckError(err)
	err = services.UnregisterService(service)
	utils.CheckError(err)
	s := "success"
	result = &grpc.Result{Result: &s}
	return
}

func (server *Server) RegisterPermission(ctx context.Context, in *grpc.TokenWithPermission) (result *grpc.Result, err error) {
	defer recoverError()
	service, err := tokens.FindServiceByTokenStr(*in.Token)
	utils.CheckError(err)
	_, err = permissions.RegisterPermission(service, *in.Permission)
	utils.CheckError(err)
	s := "success"
	result = &grpc.Result{Result: &s}
	return
}

func (server *Server) UnregisterPermission(ctx context.Context, in *grpc.TokenWithPermission) (result *grpc.Result, err error) {
	defer recoverError()
	service, err := tokens.FindServiceByTokenStr(*in.Token)
	utils.CheckError(err)
	servicePermission, err := permissions.FindServicePermissionByPermission(service, *in.Permission)
	utils.CheckError(err)
	err = permissions.UnregisterPermission(service, servicePermission)
	utils.CheckError(err)
	s := "success"
	result = &grpc.Result{Result: &s}
	return
}

func (server *Server) Authenticate(ctx context.Context, in *grpc.AuthenticateParams) (result *grpc.Result, err error) {
	defer recoverError()
	userToken, err := tokens.FindUserTokenByTokenStr(*in.UserToken)
	utils.CheckError(err)
	service, err := tokens.FindServiceByTokenStr(*in.ServiceToken)
	utils.CheckError(err)
	if !userToken.BelongsToService(service) {
		err := errors.New("invalid_token")
		utils.CheckError(err)
	}
	user, err := tokens.FindUserByTokenStr(*in.UserToken)
	utils.CheckError(err)
	if *in.Permission != "" {
		_, err = permissions.FindUserPermissionByPermission(user, *in.Permission)
		utils.CheckError(err)
	}
	result = &grpc.Result{Result: &user.UUID}
	return
}

func (server *Server) Authorize(ctx context.Context, in *grpc.AuthorizeParams) (result *grpc.Result, err error) {
	defer recoverError()
	// give permission to user
	service, err := tokens.FindServiceByTokenStr(*in.Token)
	utils.CheckError(err)
	servicePermission, err := permissions.FindServicePermissionByPermission(service, *in.Permission)
	utils.CheckError(err)
	user, err := users.FindUserByUUID(*in.UserUUID)
	utils.CheckError(err)
	err = permissions.GrantPermission(user, servicePermission)
	utils.CheckError(err)
	s := "success"
	result = &grpc.Result{Result: &s}
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