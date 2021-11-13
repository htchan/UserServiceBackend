package tokens

import (
	"time"
	// "errors"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/users"
)

type UserToken struct {
	Token string
	Username string
	generateDate int64
	duration int
}

func generateUserToken(username string, duration int) *UserToken {
	userToken := new(UserToken)
	userToken.Token = utils.RandomString(64)
	userToken.Username = username
	userToken.generateDate = time.Now().Unix()
	userToken.duration = duration
	return userToken
}

func LoadUserToken(user users.User, duration int) (*UserToken, error) {
	// if no token exist then generate a alphanumeric string between TOKEN_MIN_LEN and TOKEN_MAX_LEN
	// if token exist and not expired then return the token
	userToken, err := FindUserTokenByUser(user)
	if err == nil {
		if time.Now().Before(
			time.Unix(userToken.generateDate, 0).Add(time.Duration(userToken.duration) * time.Minute)) {
			return userToken, nil
		}
		// if token is expired then delete old token and generate a new one
		userToken.delete()
	}
	userToken = generateUserToken(user.Username, duration)
	err = userToken.create()
	if err != nil {
		return nil, err
	} else {
		return userToken, nil
	}
}

func DeleteUserTokens(user users.User) error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from user_tokens where username=?", user.Username)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (token *UserToken) Expire() error{
	return token.delete()
}