package user

import (
	"github.com/google/uuid"
)

type User struct {
	UUID string
	Username string
	EncryptedPassword string
}

var emptyUser = User{}

func ValidUser(username, password string) error {
	if len(username) == 0 || len(password) == 0 {
		return InvalidParamsError("username or password")
	}
	return nil
}

func newUser(username, password string) (User, error) {
	if err := ValidUser(username, password); err != nil {
		return User{}, err
	}
	encryptedPassword, err := hashPassword(password)
	if err != nil { return User{}, err }
	user := User{
		UUID: uuid.NewString(),
		Username: username,
		EncryptedPassword: encryptedPassword,
	}
	return user, nil
}

func Signup(username, password string) (User, error) {
	user, err := FindUserByName(username)
	if err == nil { return User{}, DuplicatedUserError{} }
	user, err = newUser(username, password)
	if err != nil { return User{}, err }
	err = user.create()
	return user, err
}

func Dropout(user User) error {
	return user.delete()
}

func Login(username, password string) (User, error) {
	user, err := FindUserByName(username)
	if err != nil {
		return User{}, err
	}
	if checkPasswordHash(password, user.EncryptedPassword) {
		return user, nil
	} else {
		return User{}, IncorrectParamsError("username or password")
	}
}