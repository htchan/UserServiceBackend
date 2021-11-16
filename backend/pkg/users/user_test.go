package users

import (
	"os"
	"io"
	"testing"
	"github.com/htchan/UserService/backend/internal/utils"
)

func init() {
	// copy database to test environment
	source, err := os.Open("../../assets/template.db")
	utils.CheckError(err)
	destination, err := os.Create("../../test/users/user-test-data.db")
	utils.CheckError(err)
	io.Copy(destination, source)
	source.Close()
	destination.Close()
}

func Test_newUser(t *testing.T) {
	utils.OpenDB("../../test/users/user-test-data.db")
	defer utils.CloseDB()

	t.Run("success", func(t *testing.T) {
		user, err := newUser("username", "password")
		if user == nil || err != nil || user.Username != "username" ||
			!checkPasswordHash("password", user.encryptedPassword) ||
			user.UUID == "" {
			t.Fatalf("user.newUser(\"username\", \"password\") return user: %v, err: %v",
				user, err)
		}
	})
}

func TestSignup(t *testing.T) {
	utils.OpenDB("../../test/users/user-test-data.db")
	defer utils.CloseDB()

	t.Run("success", func(t *testing.T) {
		user, err := Signup("signup_user", "password")
		if user == nil || err != nil ||
			user.Username != "signup_user" || user.UUID == "" {
			t.Fatalf("user.SignUp(\"signup_user\", \"password\") return user: %v, err: %v",
				user, err)
		}
		user, err = FindUserByName("signup_user")
		if user == nil || err != nil {
			t.Fatalf("user.SignUp(\"signup_user\", \"password\") saved record cannot be retrieved")
		}
	})

	t.Run("already taken username", func(t *testing.T) {
		Signup("taken_user", "password")
		user, err := Signup("taken_user", "password")
		if user != nil || err == nil {
			t.Fatalf("user.Signup allow duplicated username")
		}
	})
}

func TestLogin(t *testing.T) {
	utils.OpenDB("../../test/users/user-test-data.db")
	defer utils.CloseDB()

	_, err := Signup("login_user", "password")
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		user, err := Login("login_user", "password")
		if user == nil || err != nil {
			t.Fatalf("user.Login(\"login_user\", \"password\") return user: %v, err: %v",
				user, err)
		}
	})
	
	t.Run("wrong password", func(t *testing.T) {
		user, err := Login("login_user", "wrong_password")
		if user != nil || err == nil {
			t.Fatalf("user.Login(\"login_user\", \"wrong_password\") login successfully")
		}
	})
	
	t.Run("not exist user", func(t *testing.T) {
		user, err := Login("not_exist_user", "password")
		if user != nil || err == nil {
			t.Fatalf("user.Login(\"not_exist_user\", \"password\") login successfully")
		}
	})
}
