package token

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/service"
	"github.com/htchan/UserService/backend/pkg/user"
	"time"
)

type UserToken struct {
	UserUUID     string
	ServiceUUID  string
	Token        string
	generateDate time.Time
	duration     time.Duration
}

const (
	UserTokenLength = 64
	defaultDuration = 30 * 24 * time.Hour
)

var emptyUserToken = UserToken{}
var TokenExpiredError = errors.New("token Already Expired")

func generateUserToken() string {
	b := make([]byte, UserTokenLength)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)[:UserTokenLength]
}

func NewUserToken(u user.User, s service.Service) UserToken {
	return UserToken{
		UserUUID:     u.UUID,
		ServiceUUID:  s.UUID,
		Token:        generateUserToken(),
		generateDate: time.Now(),
		duration:     defaultDuration,
	}
}

func GetUserToken(tokenString string) (UserToken, error) {
	return queryUserToken(
		UserToken{Token: tokenString}, "select user token",
		"select user_uuid, service_uuid, token, created_date, duration from user_tokens where token=?",
		tokenString,
	)
}

func DeleteAllUserTokens(u user.User) error {
	return utils.Execute(
		UserToken{UserUUID: u.UUID}, "delete all user token",
		"delete from user_tokens where user_uuid=?",
		u.UUID,
	)
}

func DeleteAllExpiredUserTokens(u user.User) error {
	for token := range queryUserTokens(
		UserToken{UserUUID: u.UUID}, "select user tokens",
		"select user_uuid, service_uuid, token, created_date, duration from user_tokens where user_uuid=?",
		u.UUID,
	) {
		if err := token.Valid(); err != nil {
			token.Delete()
		}
	}
	return nil
}

func (t UserToken) Create() error {
	return utils.Execute(
		t, "create user token",
		"insert into user_tokens (user_uuid, service_uuid, token, created_date, duration) values (?, ?, ?, ?, ?)",
		t.UserUUID, t.ServiceUUID, t.Token, t.generateDate, t.duration,
	)
}

func (t UserToken) Delete() error {
	return utils.Execute(
		t, "delete user token",
		"delete from user_tokens where token=?",
		t.Token,
	)
}

func (t *UserToken) Expire() error {
	t.duration = -1
	return utils.Execute(
		*t, "force expire user token",
		"update user_tokens set duration=? where token=?",
		t.duration, t.Token,
	)
}

func (t UserToken) Valid() error {
	if time.Now().After(t.generateDate.Add(t.duration)) {
		return TokenExpiredError
	}
	return nil
}

func (t UserToken) User() (user.User, error) {
	return user.GetUser(t.UserUUID)
}

func (t UserToken) Service() (service.Service, error) {
	return service.GetService(t.ServiceUUID)
}
