package tokens

import (
	"time"
	// "errors"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/users"
	"github.com/htchan/UserService/backend/pkg/services"
)

const ExpireMinutes = 24 * 60

type UserToken struct {
	Token string
	userUUID string
	serviceUUID string
	generateDate int64
	duration int
}

func GenerateUserToken(user *users.User, service *services.Service, duration int) (*UserToken, error) {
	if duration < 0 { duration = ExpireMinutes }
	userToken := new(UserToken)
	for true {
		userToken.Token = utils.RandomString(64)
		if _, err := FindUserTokenByTokenStr(userToken.Token); err != nil {
			break
		}
	}
	userToken.userUUID = user.UUID
	userToken.serviceUUID = service.UUID
	userToken.generateDate = time.Now().Unix()
	userToken.duration = duration
	err := userToken.create()
	return userToken, err
}

func LoadUserToken(user *users.User, service *services.Service, duration int) (*UserToken, error) {
	// if no token exist then generate a alphanumeric string between TOKEN_MIN_LEN and TOKEN_MAX_LEN
	// if token exist and not expired then return the token
	userToken, err := FindUserTokenByUserService(user, service)
	if err == nil {
		if time.Now().Before(
			time.Unix(userToken.generateDate, 0).Add(time.Duration(userToken.duration) * time.Minute)) {
			return userToken, nil
		}
		// if token is expired then delete old token and generate a new one
		err = userToken.delete()
		if err != nil {
			return nil, err
		}
	}
	return GenerateUserToken(user, service, duration)
}

func DeleteUserTokens(user *users.User) error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from user_tokens where user_uuid=?", user.UUID)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (token *UserToken) Expire() error{
	return token.delete()
}

func (token *UserToken) BelongsToService(service *services.Service) bool {
	return token.serviceUUID == service.UUID
}