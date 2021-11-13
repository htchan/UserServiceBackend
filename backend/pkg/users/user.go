package users

import (
	"errors"
	// "github.com/htchan/UserService/backend/internal/utils"
)

type User struct {
	Username string
	encryptedPassword string
	CreatedAt int
	UpdatedAt int
}

func newUser(username, password string) (*User, error) {
	user := new(User)
	var err error
	user.Username = username
	user.encryptedPassword, err = hashPassword(password)
	return user, err
}

func Signup(username, password string) (*User, error) {
	user, err := FindUserByName(username)
	if err == nil {
		return nil, errors.New("username already been taken")
	}
	user, err = newUser(username, password)
	if err != nil {
		return nil, err
	}
	err = user.create()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func Dropout(user User) error {
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
		return nil, errors.New("incorrect username or password")
	}
}
