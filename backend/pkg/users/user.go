package users

import (
	"errors"
	"github.com/google/uuid"
	// "github.com/htchan/UserService/backend/internal/utils"
)

type User struct {
	UUID string
	Username string
	encryptedPassword string
}

func check(username, password string) bool {
	return len(username) != 0 && len(password) != 0
}

func newUser(username, password string) (*User, error) {
	if !check(username, password) {
		return nil, errors.New("invalid_username_password")
	}
	user := new(User)
	user.UUID = uuid.NewString()
	user.Username = username
	user.encryptedPassword, _ = hashPassword(password)
	return user, nil
}

func Signup(username, password string) (*User, error) {
	user, err := FindUserByName(username)
	if err == nil {
		return nil, errors.New("username_already_exist")
	}
	user, err = newUser(username, password)
	if err != nil {
		return nil, err
	}
	err = user.create()
	return user, err
}

func Dropout(user *User) error {
	return user.delete()
}

func Login(username, password string) (*User, error) {
	user, err := FindUserByName(username)
	if err != nil {
		return nil, err
	}
	if checkPasswordHash(password, user.encryptedPassword) {
		return user, nil
	} else {
		return nil, errors.New("incorrect_username_password")
	}
}
