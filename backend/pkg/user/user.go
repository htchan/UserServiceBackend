package user

import (
	"errors"
	"github.com/google/uuid"
	"github.com/htchan/UserService/backend/internal/utils"
)

type User struct {
	UUID              string
	Username          string
	EncryptedPassword string
}

var emptyUser = User{}

func NewUser(username, password string) User {
	encryptedPassword, err := hashPassword(password)
	if err != nil {
		return emptyUser
	}
	return User{
		UUID:              uuid.NewString(),
		Username:          username,
		EncryptedPassword: encryptedPassword,
	}
}

func (u User) Create() error {
	return utils.Execute(
		u, "create user",
		"insert into users (uuid, username, password) values (?, ?, ?)",
		u.UUID, u.Username, u.EncryptedPassword,
	)
}

func (u User) Delete() error {
	return utils.Execute(
		u, "delete user",
		"delete from users where UUID=?",
		u.UUID,
	)
}


func GetUser(uuid string) (User, error) {
	return queryUser(
		User{UUID: uuid}, "query user by uuid",
		"select uuid, username, password from users where uuid=?",
		uuid,
	)
}

func GetUserByName(username, password string) (User, error) {
	u := User{Username: username}
	operation := "query user by username"
	user, err := queryUser(
		u, operation,
		"select uuid, username, password from users where username=?",
		username,
	)
	if err != nil {
		return u, err
	}
	if !checkPasswordHash(password, user.EncryptedPassword) {
		return emptyUser, utils.NewNotFoundError(operation, u, errors.New("wrong password"))
	}
	return u, nil
}

func (u User) Valid(password string) error {
	if !checkPasswordHash(password, u.EncryptedPassword) {
		return utils.NewInvalidRecordError("wrong password")
	}
	return nil
}