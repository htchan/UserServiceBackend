package grpc

import (
	"google.golang.org/grpc/test/bufconn"
	"google.golang.org/grpc"
	pb "github.com/htchan/UserService/backend/internal/grpc"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/users"
	"github.com/htchan/UserService/backend/pkg/tokens"
	"github.com/htchan/UserService/backend/pkg/services"
	"github.com/htchan/UserService/backend/pkg/permissions"

	"context"
	"net"
	"testing"
	"log"
	"io"
	"os"
)

const bufSize = 1024 * 1024

var lis *bufconn.Listener

var ctx context.Context
var conn *grpc.ClientConn

var client pb.UserServiceClient

func init() {
	source, err := os.Open("../../assets/template.db")
	utils.CheckError(err)
	destination, err := os.Create("../../test/grpc/server-test-data.db")
	utils.CheckError(err)
	io.Copy(destination, source)
	source.Close()
	destination.Close()
    lis = bufconn.Listen(bufSize)
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &Server{})
    go func() {
		log.Fatal(s.Serve(lis))
    }()
    ctx = context.Background()
    conn, err = grpc.DialContext(ctx, "bufnet", grpc.WithContextDialer(bufDialer), grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Failed to dial bufnet: %v", err)
    }
    client = pb.NewUserServiceClient(conn)
}

func bufDialer(context.Context, string) (net.Conn, error) {
    return lis.Dial()
}

func TestSignup(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()

	username := "signup_username"
	password := "password"
	loginParams := NewLoginParams(username, password)
	
	t.Run("success", func(t *testing.T) {
		token, err := client.Signup(ctx, loginParams)
		if token == nil || err != nil || len(*token.Token) != 64 {
			t.Fatalf("grpc failed in normal signup: token - %v, err - %v",
				token, err)
		}
		user, err := users.FindUserByName(username)
		if user == nil || err != nil || user.Username != username {
			t.Fatalf("grpc.Server.Signup cannot save user")
		}
	})

	t.Run("fail if username already exist", func(t *testing.T) {
		token, err := client.Signup(ctx, loginParams)
		if token != nil || err == nil {
			t.Fatalf("grpc failed in repeat username signup: token - %v, err - %v",
				token, err)
		}
	})
}
func TestDropout(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()

	username := "dropout_username"
	password := "password"
	loginParams := NewLoginParams(username, password)
	token, err := client.Signup(ctx, loginParams)
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		result, err := client.Dropout(ctx, token)
		if result == nil || err != nil || *result.Result != "success" {
			t.Fatalf("grpc.Server.Dropout fail to remove user: result - %v, err - %v",
				result, err)
		}
		user, err := users.FindUserByName(username)
		if user != nil || err == nil {
			t.Fatalf("ggrpc fail to remove user: user - %v, err - %v",
				user, err)
		}
	})

	t.Run("fail if token cannot map any user", func(t *testing.T) {
		result, err := client.Dropout(ctx, token)
		if result != nil || err == nil {
			t.Fatalf("grpc.Server.Dropout fail on not exist token: result - %v, err - %v",
				result, err)
		}
	})
}

func TestLogin(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()

	username := "login_username"
	password := "password"
	invalidPassword := "invalid password"
	loginParams := NewLoginParams(username, password)
	token, err := client.Signup(ctx, loginParams)
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		actualToken, err := client.Login(ctx, loginParams)
		if actualToken == nil || err != nil || *actualToken.Token != *token.Token {
			t.Fatalf("grpc.Server.Login fail in normal flow: token - %v, err - %v",
				actualToken, err)
		}
	})

	t.Run("success even token in db expired", func(t *testing.T) {
		dbToken, err := tokens.FindUserTokenByTokenStr(*token.Token)
		utils.CheckError(err)
		dbToken.Expire()
		actualToken, err := client.Login(ctx, loginParams)
		if actualToken == nil || err != nil || *actualToken.Token == *token.Token {
			t.Fatalf("grpc.Server.Login fail in normal flow: token - %v, err - %v",
				actualToken, err)
		}
		token = actualToken
	})
	
	t.Run("fail if password invalid", func(t *testing.T) {
		invalidLoginParams := NewLoginParams(username, invalidPassword)
		actualToken, err := client.Login(ctx, invalidLoginParams)
		if actualToken != nil || err == nil {
			t.Fatalf("grpc.Server.Login success for invalid password: token - %v, err - %v",
				actualToken, err)
		}
	})
}
func TestLogout(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()

	username := "logout_username"
	password := "password"
	loginParams := NewLoginParams(username, password)
	token, err := client.Signup(ctx, loginParams)
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		result, err := client.Logout(ctx, token)
		if result == nil || err != nil || *result.Result != "success" {
			t.Fatalf("grpc.Server.Logout failed in normal flow: result - %v, err - %v",
				result, err)
		}
	})

	t.Run("fail if user already logout", func(t *testing.T) {
		result, err := client.Logout(ctx, token)
		if result != nil || err == nil {
			t.Fatalf("grpc.Server.Logout success for logout user: result - %v, err - %v",
				result, err)
		}
	})
}

func TestRegisterService(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()

	name := "reg_service"
	serviceName := NewServiceName(name)

	t.Run("success", func(t *testing.T) {
		token, err := client.RegisterService(ctx, serviceName)
		if token == nil || err != nil || len(*token.Token) != 64 {
			t.Fatalf("grpc.Server.RegisterService fail in normal flow: token - %v, err - %v",
				token, err)
		}
		service, err := services.FindServiceByName(name)
		if service == nil || err != nil {
			t.Fatalf("grpc.Server.RegisterService does not save service to db")
		}
		actualToken, err := tokens.FindServiceTokenByTokenStr(*token.Token)
		if actualToken == nil || err != nil {
			t.Fatalf("grpc.Server.RegisterService does not save token to db")
		}
 	})

	t.Run("fail for already registered service name", func(t *testing.T) {
		token, err := client.RegisterService(ctx, serviceName)
		if token != nil || err == nil {
			t.Fatalf("grpc.Server.RegisterService success for repeated service name: token - %v, err - %v",
				token, err)
		}
	})
}
func TestUnregisterService(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()

	name := "un_reg_service"
	serviceName := NewServiceName(name)
	token, err := client.RegisterService(ctx, serviceName)
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		result, err := client.UnregisterService(ctx, token)
		if result == nil || err != nil || *result.Result != "success" {
			t.Fatalf("grpc.Server.UnregisterService fail for normal flow: result - %v, err - %v",
				result, err)
		}
		service, err := services.FindServiceByName(name)
		if service != nil || err == nil {
			t.Fatalf("grpc.Server.UnregisterService does not remove service in db")
		}
	})

	t.Run("fail for not exist service name", func(t *testing.T) {
		result, err := client.UnregisterService(ctx, token)
		if result != nil || err == nil {
			t.Fatalf("grpc.Server.UnregisterService success for not existed service name - %v, err - %v",
				result, err)
		}
	})
}

func TestRegisterPermission(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()

	name := "reg_permission_service"
	serviceName := NewServiceName(name)
	token, err := client.RegisterService(ctx, serviceName)
	utils.CheckError(err)

	permissionName := "reg_permission"
	permissionWithToken := NewTokenWithPermission(*token.Token, permissionName)

	t.Run("success", func(t *testing.T) {
		result, err := client.RegisterPermission(ctx, permissionWithToken)
		if result == nil || err != nil || *result.Result != "success" {
			t.Fatalf("grpc.Server.RegisterPermission fail in normal flow: result - %v, err - %v",
				result, err)
		}
		service, err := services.FindServiceByName(name)
		utils.CheckError(err)
		permission, err := permissions.FindServicePermissionByPermission(service, permissionName)
		if permission == nil || err != nil {
			t.Fatalf("grpc.Server.RegisterPermission does not save service permission to db")
		}
	})

	t.Run("fail if permission already exist", func(t *testing.T) {
		result, err := client.RegisterPermission(ctx, permissionWithToken)
		if result != nil || err == nil {
			t.Fatalf("grpc.Server.RegisterPermission success for existing service permission: result - %v, err - %v",
				result, err)
		}
	})
}
func TestUnregisterPermission(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()
	
	name := "un_reg_permission_service"
	serviceName := NewServiceName(name)
	token, err := client.RegisterService(ctx, serviceName)
	utils.CheckError(err)

	emptyName := "empty_permission_service"
	emptyServiceName := NewServiceName(emptyName)
	emptyServiceToken, err := client.RegisterService(ctx, emptyServiceName)
	utils.CheckError(err)

	permissionName := "un_reg_permission"
	permissionWithToken := NewTokenWithPermission(*token.Token, permissionName)
	_, err = client.RegisterPermission(ctx, permissionWithToken)
	utils.CheckError(err)

	t.Run("faul if permission is not own by service", func(t *testing.T) {
		emptyPermissionWithToken := NewTokenWithPermission(*emptyServiceToken.Token, permissionName)
		result, err := client.UnregisterPermission(ctx, emptyPermissionWithToken)
		if result != nil || err == nil {
			t.Fatalf("grpc.Server.UnregisterPermission success in not match service")
		}

		service, err := services.FindServiceByName(name)
		utils.CheckError(err)
		permission, err := permissions.FindServicePermissionByPermission(service, permissionName)
		if permission == nil || err != nil {
			t.Fatalf("grpc.Server.UnregisterPermission of wrong service remove permission from db")
		}
	})

	t.Run("success", func(t *testing.T) {
		result, err := client.UnregisterPermission(ctx, permissionWithToken)
		if result == nil || err != nil || *result.Result != "success" {
			t.Fatalf("grpc.Server.UnregisterPermission fail in normal flow: result - %v, err - %v",
				result, err)
		}

		service, err := services.FindServiceByName(name)
		utils.CheckError(err)
		permission, err := permissions.FindServicePermissionByPermission(service, permissionName)
		if permission != nil || err == nil {
			t.Fatalf("grpc.Server.UnregisterPermission does not remove permission from db")
		}
	})

	t.Run("fail for not exist service", func(t *testing.T) {
		result, err := client.UnregisterPermission(ctx, permissionWithToken)
		if result != nil || err == nil {
			t.Fatalf("grpc.Server.UnregisterPermission success for not exist permission")
		}
	})
}

func TestAuthorize(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()
	name := "author_service"
	serviceName := NewServiceName(name)
	serviceToken, err := client.RegisterService(ctx, serviceName)
	utils.CheckError(err)
	permissionName := "author_permission"
	permissionWithToken := NewTokenWithPermission(*serviceToken.Token, permissionName)
	_, err = client.RegisterPermission(ctx, permissionWithToken)
	utils.CheckError(err)

	username := "author_username"
	password := "password"
	token, err := client.Signup(ctx, NewLoginParams(username, password))
	utils.CheckError(err)

	user, err := tokens.FindUserByTokenStr(*token.Token)
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		result, err := client.Authorize(ctx, NewAuthorizeParams(
			*serviceToken.Token,
			user.UUID,
			permissionName,
		))
		if result == nil || err != nil || *result.Result != "success" {
			t.Fatalf("grpc.Server.Authorize fail in normal flow: result - %v, err - %v",
				result, err)
		}

		user, err := users.FindUserByName(username)
		utils.CheckError(err)
		userPermission, err := permissions.FindUserPermissionByPermission(user, permissionName)
		if userPermission == nil || err != nil {
			t.Fatalf("grpc.Server.Authorize does not save user permission to db")
		}
	})

	t.Run("fail if user already authenticated", func(t *testing.T) {
		result, err := client.Authorize(ctx, NewAuthorizeParams(
			*serviceToken.Token,
			user.UUID,
			permissionName,
		))
		if result != nil || err == nil {
			t.Fatalf("grpc.Server.Authorize success for existing user permission: result - %v, err - %v",
				result, err)
		}
	})
}
func TestAuthenticate(t *testing.T) {
	utils.OpenDB("../../test/grpc/server-test-data.db")
	defer utils.CloseDB()
	name := "authen_service"
	serviceName := NewServiceName(name)
	serviceToken, err := client.RegisterService(ctx, serviceName)
	utils.CheckError(err)
	permissionName := "authen_permission"
	permissionWithToken := NewTokenWithPermission(*serviceToken.Token, permissionName)
	_, err = client.RegisterPermission(ctx, permissionWithToken)
	utils.CheckError(err)

	username := "authen_username"
	password := "password"
	userToken, err := client.Signup(ctx, NewLoginParams(username, password))
	utils.CheckError(err)
	user, err := tokens.FindUserByTokenStr(*userToken.Token)
	utils.CheckError(err)

	username2 := "no_authen_username"
	password2 := "password"
	userToken2, err := client.Signup(ctx, NewLoginParams(username2, password2))
	utils.CheckError(err)
	_, err = client.Authorize(ctx, NewAuthorizeParams(
		*serviceToken.Token,
		user.UUID,
		permissionName,
	))
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		result, err := client.Authenticate(ctx, NewTokenWithPermission(
			*userToken.Token,
			permissionName,
		))
		if result == nil || err != nil || *result.Result != "success" {
			t.Fatalf("grpc.Server.Authenticate fail in normal flow")
		}
	})

	t.Run("fail if user have no permission", func(t *testing.T) {
		result, err := client.Authenticate(ctx, NewTokenWithPermission(
			*userToken2.Token,
			permissionName,
		))
		if result != nil || err == nil {
			t.Fatalf("grpc.Server.Authenticate success for user have no permission")
		}
	})
}