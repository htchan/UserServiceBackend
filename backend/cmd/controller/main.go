package main


import (
	"flag"
	"fmt"
	"strings"
	"context"
	"github.com/htchan/UserService/backend/pkg/grpc"
)

var (
	addr *string
	operation *string
	username *string
	password *string
	serviceName *string
	url *string
	token *string
	permission *string
	userUUID *string
	client grpc.Client
	ctx context.Context
)

func initFlag() {
	addr = flag.String("addr", "", "address of grpc server")
	operation = flag.String("operation", "", "operation to carry")
	username = flag.String("username", "", "username")
	password = flag.String("password", "", "password")
	serviceName = flag.String("service-name", "", "service name")
	url = flag.String("url", "", "url")
	token = flag.String("token", "", "user / service token")
	permission = flag.String("permission", "", "permission")
	userUUID = flag.String("user-uuid", "", "user uuid")
	flag.Parse()
}

func checkVariableNotEmpty(name string, s *string) {
	if *s == "" {
		panic(fmt.Sprintf("parameter not provided: %v", name))
	}
}

func reportError(operation string, err error) {
	if err != nil {
		panic(fmt.Sprintf("%v error: %v", operation, err))
	}
}

func Signup() {
	checkVariableNotEmpty("username", username)
	checkVariableNotEmpty("password", password)

	loginParams := grpc.NewLoginParams(*username, *password)
	token, err := client.Signup(ctx, loginParams)

	reportError("signup", err)
	fmt.Printf("token: %v", token)
}

func Dropout() {
	checkVariableNotEmpty("token", token)

	authToken := grpc.NewAuthToken(*token)
	result, err := client.Dropout(ctx, authToken)

	reportError("dropout", err)
	fmt.Printf("result: %v", result)
}

func Login() {
	checkVariableNotEmpty("username", username)
	checkVariableNotEmpty("password", password)

	loginParams := grpc.NewLoginParams(*username, *password)
	authToken, err := client.Login(ctx, loginParams)

	reportError("login", err)
	fmt.Printf("token: %v", authToken)
}

func Logout() {
	checkVariableNotEmpty("token", token)

	authToken := grpc.NewAuthToken(*token)
	result, err := client.Logout(ctx, authToken)

	reportError("logout", err)
	fmt.Printf("result: %v", result)
}

func RegisterService() {
	checkVariableNotEmpty("service-name", serviceName)
	checkVariableNotEmpty("url", url)

	service := grpc.NewServiceName(*serviceName, *url)
	token, err := client.RegisterService(ctx, service)

	reportError("register service", err)
	fmt.Printf("token: %v", token)
}

func UnregisterService() {
	checkVariableNotEmpty("token", token)

	authToken := grpc.NewAuthToken(*token)
	result, err := client.UnregisterService(ctx, authToken)

	reportError("logout", err)
	fmt.Printf("result: %v", result)
}

func Authenticate() {
	checkVariableNotEmpty("token", token)
	checkVariableNotEmpty("permission", permission)

	tokenPermission := grpc.NewTokenWithPermission(*token, *permission)
	result, err := client.Authenticate(ctx, tokenPermission)

	reportError("authenticate", err)
	fmt.Printf("result: %v", result)
}

func Authorize() {
	checkVariableNotEmpty("token", token)
	checkVariableNotEmpty("user-uuid", userUUID)
	checkVariableNotEmpty("permission", permission)

	authorizeParams := grpc.NewAuthorizeParams(*token, *userUUID, *permission)
	result, err := client.Authorize(ctx, authorizeParams)

	reportError("authorize", err)
	fmt.Printf("result: %v", result)
}

func RegisterPermission() {
	checkVariableNotEmpty("token", token)
	checkVariableNotEmpty("permission", permission)

	tokenPermission := grpc.NewTokenWithPermission(*token, *permission)
	result, err := client.RegisterPermission(ctx, tokenPermission)

	reportError("register permission", err)
	fmt.Printf("result: %v", result)
}

func UnregisterPermission() {
	checkVariableNotEmpty("token", token)
	checkVariableNotEmpty("permission", permission)

	tokenPermission := grpc.NewTokenWithPermission(*token, *permission)
	result, err := client.UnregisterPermission(ctx, tokenPermission)

	reportError("unregister permission", err)
	fmt.Printf("result: %v", result)
}

func main() {
	initFlag()
	checkVariableNotEmpty("addr", addr)
	funcMap := map[string]func() {
		"signup": Signup,
		"dropout": Dropout,
		"login": Login,
		"logout": Logout,
		"register-service": RegisterService,
		"unregister-service": UnregisterService,
		"authenticate": Authenticate,
		"authorize": Authorize,
		"register-permission": RegisterPermission,
		"unregister-permission": UnregisterPermission,
	}
	*operation = strings.ToLower(*operation)
	if _, ok := funcMap[*operation]; !ok {
		panic(fmt.Sprintf("operation not found: %v", *operation))
	}
	ctx = context.Background()
	client = grpc.NewClient(*addr)
	funcMap[*operation]()
}