package tokens

import (
	"testing"
	"os"
	"io"
	"time"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/users"
)

func init() {
	// copy database to test environment
	source, err := os.Open("../../assets/template.db")
	utils.CheckError(err)
	destination, err := os.Create("../../test/tokens/user-token-test-data.db")
	utils.CheckError(err)
	io.Copy(destination, source)
	source.Close()
	destination.Close()
}

func Test_generateToken(t *testing.T) {
	utils.OpenDB("../../test/tokens/user-token-test-data.db")
	defer utils.CloseDB()

	user, err := users.Signup("username", "password")
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		token := generateUserToken(user, 100)
		if token.Token == "" || token.userUUID != user.UUID ||
			token.duration != 100 {
			t.Fatalf("tokens.generateUserToken return wrong token in normal flow")
		}
	})
}

func TestLoadUserToken(t *testing.T) {
	utils.OpenDB("../../test/tokens/user-token-test-data.db")
	defer utils.CloseDB()
	t.Run("user already have token", func(t *testing.T) {
		user, err := users.Signup("token_owner", "password")
		utils.CheckError(err)
		token := generateUserToken(user, 10)
		if token == nil {
			panic("token is null")
		}
		token.create()
		
		resultToken, err := LoadUserToken(user, 100)
		if err != nil || resultToken.Token != token.Token {
			t.Fatalf("tokens.LoadUserToken returns wrong token\nexpect: %v\nactual: %v\nerror: %v",
				token, resultToken, err)
		}
	})

	t.Run("user do not have token", func(t *testing.T) {
		user, err := users.Signup("no_token_owner", "password")
		utils.CheckError(err)
		now := time.Now()
		resultToken, err := LoadUserToken(user, 100)
		if err != nil || resultToken.generateDate < now.Unix() {
			t.Fatalf("tokens.LoadUserToken returns old generated token: %v, error: %v",
				resultToken.Token, err)
		}
	})

	t.Run("user's token already expired", func(t *testing.T) {
		user, err := users.Signup("expire_token_owner", "password")
		utils.CheckError(err)
		token, err := LoadUserToken(user, 0)
		utils.CheckError(err)
		now := time.Now()
		
		resultToken, err := LoadUserToken(user, 100)
		if err != nil || resultToken.generateDate < now.Unix() || resultToken.Token == token.Token {
			t.Fatalf("tokens.LoadUserToken returns old generated token: %v, error: %v",
				resultToken, err)
		}
		checkToken, err := FindUserTokenByTokenStr(token.Token)
		if checkToken != nil || err == nil {
			t.Fatalf("tokens.LoadUserToken loads expired token")
		}
	})
}

func TestDeleteUserTokens(t *testing.T) {
	utils.OpenDB("../../test/tokens/user-token-test-data.db")
	defer utils.CloseDB()
	user, err := users.Signup("DeleteUsername", "password")
	token := generateUserToken(user, 100)
	err = token.create()
	utils.CheckError(err)
	emptyUser, err := users.Signup("empty_user", "password")
	utils.CheckError(err)
	

	t.Run("success", func(t *testing.T) {
		err = DeleteUserTokens(user)
		utils.CheckError(err)
		token, err = FindUserTokenByUser(user)
		if token != nil || err == nil {
			t.Fatalf("tokens.Token.DeleteUserToken() cannot delete token: %v, err: %v",
				token, err)
		}
	})

	t.Run("success for user with no token", func(t *testing.T) {
		err = DeleteUserTokens(emptyUser)
		if err != nil {
			t.Fatalf("tokens.DeleteUserToken of no token raise error %v", err)
		}
	})
}
