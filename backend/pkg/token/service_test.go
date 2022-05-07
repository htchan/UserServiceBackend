package token

import (
	"testing"
	"errors"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/user"
	"github.com/htchan/UserService/backend/pkg/service"
)

func TestUserSignUp(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()

	t.Run("signup valid user", func (t *testing.T) {
		tkn, err := UserSignup("user signup", "signup")
		s, _ := tkn.Service()
		if err != nil || s.UUID != service.DefaultUserService().UUID {
			t.Errorf("return error - tkn: %v, err: %v", tkn, err)
			t.Errorf("%v, %v", s.UUID, service.DefaultUserService().UUID)
		}
	})

	t.Run("signup user already exist", func (t *testing.T) {
		tkn, err := UserSignup("user signup", "signup")
		if !errors.Is(err, utils.DatabaseError) || tkn != emptyUserToken {
			t.Errorf("not return error - tkn: %v, err: %v", tkn, err)
		}
	})
}

func TestUserDropout(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()
	
	t.Run("success", func (t *testing.T) {
		tkn, _ := UserSignup("user dropout", "dropout")
		u, _ := tkn.User()
		err := UserDropout(tkn.Token)
		if err != nil {
			t.Errorf("return error: %v", err)
		}
		u, err = user.GetUser(u.UUID)
		if !errors.Is(err, utils.NotFoundError) {
			t.Errorf("user was not deleted")
		}
		tkn, err = GetUserToken(tkn.Token)
		if !errors.Is(err, utils.NotFoundError) {
			t.Errorf("token was not deleted")
		}
	})
	
	t.Run("fail if user not exist", func (t *testing.T) {
		err := UserDropout("abc")
		if !errors.Is(err, InvalidTokenError) {
			t.Errorf("return error: %v", err)
		}
	})
	
	t.Run("fail if token is not from default user service", func (t *testing.T) {
		tkn, _ := UserSignup("user dropout", "dropout")
		s := service.NewService("dropout service", "http://localhost")
		s.Create()
		otherServiceTkn, err := UserTokenLogin(tkn.Token, s.UUID)
		u, _ := tkn.User()
		err = UserDropout(otherServiceTkn.Token)
		if !errors.Is(err, InvalidTokenError) {
			t.Errorf("return error: %v", err)
		}
		_, err = user.GetUser(u.UUID)
		if err != nil {
			t.Errorf("user was deleted")
		}
		_, err = GetUserToken(tkn.Token)
		if err != nil {
			t.Errorf("token was deleted")
		}
	})
}

func TestUserNameLogin(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()
	
	t.Run("login existing user", func (t *testing.T) {
		UserSignup("name login", "name login")
		tkn, url, err := UserNameLogin("name login", "name login", service.DefaultUserService().UUID)
		if err != nil || tkn.Token == "" || url != service.DefaultUserService().URL {
			t.Errorf("return error: %v", err)
		}
	})
	
	t.Run("login not exist user", func (t *testing.T) {
		_, _, err := UserNameLogin("not exist user", "name login", service.DefaultUserService().UUID)
		if !errors.Is(err, utils.NotFoundError) {
			t.Errorf("return error: %v", err)
		}
	})
}

func TestUserTokenLogin(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()
	
	t.Run("login existing user", func (t *testing.T) {
		tkn, err := UserSignup("token login", "token login")
		tknResult, err := UserTokenLogin(tkn.Token, service.DefaultUserService().UUID)
		if err != nil || tknResult.Token == "" {
			t.Errorf("return error: %v", err)
		}
	})
	
	t.Run("login not exist user", func (t *testing.T) {
		_, err := UserTokenLogin("abc", service.DefaultUserService().UUID)
		if !errors.Is(err, utils.NotFoundError) {
			t.Errorf("return error: %v", err)
		}
	})
	
	t.Run("login with token from other service", func (t *testing.T) {
		tkn, _ := UserSignup("other token login", "token login")
		s := service.NewService("login service", "http://localhost")
		s.Create()
		otherServiceTkn, _ := UserTokenLogin(tkn.Token, s.UUID)

		_, err := UserTokenLogin(otherServiceTkn.Token, service.DefaultUserService().UUID)
		if !errors.Is(err, InvalidTokenError) {
			t.Errorf("return error: %v", err)
		}
	})
}

func TestUserLogout(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()
	
	t.Run("logout existing user", func (t *testing.T) {
		tkn, err := UserSignup("token logout", "abc")
		err = UserLogout(tkn.Token)
		if err != nil {
			t.Errorf("return error: %v", err)
		}
	})
	
	t.Run("logout not exist user", func (t *testing.T) {
		err := UserLogout("abc")
		if !errors.Is(err, InvalidTokenError) {
			t.Errorf("return error: %v", err)
		}
	})
}

func TestServiceRegister(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()
	
	t.Run("register service", func (t *testing.T) {
		_, err := ServiceRegister("reg service", "http://localhost")
		if err != nil {
			t.Errorf("return error: %v", err)
		}
	})
	
	t.Run("register service already exist", func (t *testing.T) {
		_, err := ServiceRegister("reg service", "http://localhost")
		if !errors.Is(err, utils.DatabaseError) {
			t.Errorf("return error: %v", err)
		}
	})
}

func TestServiceUnregister(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()
	
	t.Run("unregister exist service", func (t *testing.T) {
		tkn, _ := ServiceRegister("reg service", "http://localhost")
		err := ServiceUnregister(tkn.Token)
		if err != nil {
			t.Errorf("return error: %v", err)
		}
	})
	
	t.Run("unregister not exist service", func (t *testing.T) {
		err := ServiceUnregister("abc")
		if err != nil {
			t.Errorf("return error: %v", err)
		}
	})
}