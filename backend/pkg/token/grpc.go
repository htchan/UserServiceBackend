package token

import (
	"errors"
	"github.com/htchan/UserService/backend/pkg/user"
	"github.com/htchan/UserService/backend/pkg/service"
)

var InvalidTokenError = errors.New("invalid_token")

func UserSignup(username, password string) (UserToken, error) {
	u := user.NewUser(username, password)
	if err := u.Valid(password); err != nil {
		return emptyUserToken, err
	}
	if err := u.Create(); err != nil {
		return emptyUserToken, err
	}
	tkn := NewUserToken(u, service.DefaultUserService())
	err := tkn.Create()
	return tkn, err
}

func UserDropout(token string) error {
	tkn, err := GetUserToken(token)
	if err != nil { return InvalidTokenError }
	if s, err := tkn.Service(); err != nil || s.UUID != service.DefaultUserService().UUID {
		return InvalidTokenError
	}
	u, err := tkn.User()
	if err != nil { return InvalidTokenError }
	err = DeleteAllUserTokens(u)
	if err != nil { return InvalidTokenError }

	return u.Delete()
}

func UserNameLogin(username, password, serviceUUID string) (UserToken, error) {
	u, err := user.GetUserByName(username, password)
	if err != nil { return emptyUserToken, err }
	DeleteAllExpiredUserTokens(u)
	s, err := service.GetService(serviceUUID)
	if err != nil { return emptyUserToken, err }

	tkn := NewUserToken(u, s)
	err = tkn.Create()
	return tkn, err
}

func UserTokenLogin(tokenString, serviceUUID string) (UserToken, error) {
	userTkn, err := GetUserToken(tokenString)
	if err != nil { return emptyUserToken, err }
	u, err := userTkn.User()
	if err != nil { return emptyUserToken, err }
	userTknService, err := userTkn.Service()
	if err != nil { return emptyUserToken, err }
	if userTknService.UUID != service.DefaultUserService().UUID {
		return emptyUserToken, InvalidTokenError
	}
	DeleteAllExpiredUserTokens(u)

	s, err := service.GetService(serviceUUID)
	if err != nil { return emptyUserToken, err }

	tkn := NewUserToken(u, s)
	err = tkn.Create()
	return tkn, err
}

func UserLogout(token string) error {
	tkn, err := GetUserToken(token)
	if err != nil { return InvalidTokenError }

	return tkn.Expire()
}

func ServiceRegister(name, url string) (ServiceToken, error) {
	s := service.NewService(name, url)
	err := s.Create()
	if err != nil { return emptyServiceToken, err }
	tkn := NewServiceToken(s)
	err = tkn.Create()
	return tkn, err
}

func ServiceUnregister(token string) error {
	tkn, _ := GetServiceToken(token)
	s, _ := tkn.Service()
	tkn.Delete()
	s.Delete()
	return nil
}