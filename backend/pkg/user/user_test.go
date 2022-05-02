package user

// import (
// 	"os"
// 	"io"
// 	"testing"
// 	"github.com/htchan/UserService/backend/internal/utils"
// )

// func init() {
// 	// copy database to test environment
// 	source, err := os.Open("../../assets/template.db")
// 	utils.CheckError(err)
// 	destination, err := os.Create("../../test/users/user-test-data.db")
// 	utils.CheckError(err)
// 	io.Copy(destination, source)
// 	source.Close()
// 	destination.Close()
// }

// func Test_ValidUser(t *testing.T) {
// 	t.Run("pass", func (t *testing.T) {
// 		if err := ValidUser("username", "password"); err != nil {
// 			t.Fatalf("users.ValidUser(\"username\", \"password\" return %v", err)
// 		}
// 	})

// 	t.Run("empty username", func (t *testing.T) {
// 		if err := ValidUser("", "password"); err == nil {
// 			t.Fatalf("users.ValidUser(\"\", \"password\" not return error")
// 		}
// 	})

// 	t.Run("empty password", func (t *testing.T) {
// 		if err := ValidUser("username", ""); err == nil {
// 			t.Fatalf("users.ValidUser(\"username\", \"\" not return error")
// 		}
// 	})
// }

// func Test_newUser(t *testing.T) {
// 	utils.OpenDB("../../test/users/user-test-data.db")
// 	defer utils.CloseDB()

// 	t.Run("success", func(t *testing.T) {
// 		user, err := newUser("username", "password")
// 		if err != nil || user.Username != "username" ||
// 			!checkPasswordHash("password", user.EncryptedPassword) ||
// 			user.UUID == "" {
// 			t.Fatalf("user.newUser(\"username\", \"password\") return user: %v, err: %v",
// 				user, err)
// 		}
// 	})

// 	t.Run("fail if username and password are empty", func (t *testing.T) {
// 		user, err := newUser("", "")
// 		if err == nil {
// 			t.Fatalf("user.newUser(\"\", \"\") return user: %v, err: %v",
// 				user, err)
// 		}
// 	})
// }

// func TestSignup(t *testing.T) {
// 	utils.OpenDB("../../test/users/user-test-data.db")
// 	defer utils.CloseDB()

// 	t.Run("success", func(t *testing.T) {
// 		user, err := Signup("signup_user", "password")
// 		if err != nil ||
// 			user.Username != "signup_user" || user.UUID == "" {
// 			t.Fatalf("user.SignUp(\"signup_user\", \"password\") return user: %v, err: %v",
// 				user, err)
// 		}
// 		user, err = FindUserByName("signup_user")
// 		if err != nil || user == emptyUser {
// 			t.Fatalf("user.SignUp(\"signup_user\", \"password\") saved record cannot be retrieved")
// 		}
// 	})

// 	t.Run("already taken username", func(t *testing.T) {
// 		Signup("taken_user", "password")
// 		_, err := Signup("taken_user", "password")
// 		if err == nil {
// 			t.Fatalf("user.Signup allow duplicated username")
// 		}
// 	})
// }

// func TestLogin(t *testing.T) {
// 	utils.OpenDB("../../test/users/user-test-data.db")
// 	defer utils.CloseDB()

// 	_, err := Signup("login_user", "password")
// 	utils.CheckError(err)

// 	t.Run("success", func(t *testing.T) {
// 		user, err := Login("login_user", "password")
// 		if err != nil || user == emptyUser {
// 			t.Fatalf("user.Login(\"login_user\", \"password\") return user: %v, err: %v",
// 				user, err)
// 		}
// 	})
	
// 	t.Run("wrong password", func(t *testing.T) {
// 		_, err := Login("login_user", "wrong_password")
// 		if err == nil {
// 			t.Fatalf("user.Login(\"login_user\", \"wrong_password\") login successfully")
// 		}
// 	})
	
// 	t.Run("not exist user", func(t *testing.T) {
// 		_, err := Login("not_exist_user", "password")
// 		if err == nil {
// 			t.Fatalf("user.Login(\"not_exist_user\", \"password\") login successfully")
// 		}
// 	})
// }
