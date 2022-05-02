package token

import (
	"errors"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/service"
	"github.com/htchan/UserService/backend/pkg/user"
	"io"
	"os"
	"testing"
	"time"
)

func init() {
	// copy database to test environment
	source, err := os.Open("../../assets/test_data.db")
	utils.CheckError(err)
	destination, err := os.Create("./token-test-data.db")
	utils.CheckError(err)
	io.Copy(destination, source)
	source.Close()
	destination.Close()
}

func Test_generateUserToken(t *testing.T) {
	t.Run("generate string with length 64", func(t *testing.T) {
		s := generateUserToken()
		if len(s) != UserTokenLength {
			t.Errorf("token length is %v", len(s))
		}
	})
}

func TestNewUserToken(t *testing.T) {
	u := user.NewUser("test", "test")
	s := service.NewService("test", "http://url")
	t.Run("generate user token instance", func(t *testing.T) {
		tkn := NewUserToken(u, s)
		if tkn.UserUUID != u.UUID || tkn.ServiceUUID != s.UUID ||
			len(tkn.Token) != UserTokenLength ||
			tkn.generateDate.Second() != time.Now().Second() ||
			tkn.duration != defaultDuration {
			t.Errorf("wront token instance: %v", tkn)
		}
	})
}

func TestGetUserToken(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()

	t.Run("get existing user token", func(t *testing.T) {
		tkn, err := GetUserToken("token")
		if err != nil || tkn.UserUUID != "1" || tkn.ServiceUUID != "1" ||
			tkn.Token != "token" || tkn.duration.String() != "1m0s" ||
			!tkn.generateDate.Equal(time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)) {
			t.Errorf("query wrong token - token: %v, err: %v", tkn, err)
		}
	})

	t.Run("get not exist user token", func(t *testing.T) {
		tkn, err := GetUserToken("not exist")
		if !errors.Is(err, utils.NotFoundError) {
			t.Errorf("query wrong token - token: %v, err: %v", tkn, err)
		}
	})
}

func TestDeleteAllUserTokens(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()

	t.Run("delete existing user's token", func(t *testing.T) {
		u := user.NewUser("test", "test")
		u.UUID = "1"
		err := DeleteAllUserTokens(u)
		if err != nil {
			t.Errorf("return error %v", err)
		}
		tkn, err := GetUserToken("token")
		if !errors.Is(err, utils.NotFoundError) {
			t.Errorf("user token still exist - token: %v, err: %v", tkn, err)
		}
	})
}

func TestDeleteAllExpiredUserTokens(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()

	t.Run("delete expired user's token", func(t *testing.T) {
		u := user.NewUser("expired", "test")
		u.UUID = "2"
		s := service.NewService("test", "http://url")
		tkn := NewUserToken(u, s)
		tkn.Token = "expired token"
		tkn.generateDate = time.Now().Add(-10 * time.Hour)
		tkn.duration = 1 * time.Minute
		err := tkn.Create()
		if err != nil {
			t.Errorf("faile to create token - token: %v, err: %v", tkn, err)
		}
		tkn, err = GetUserToken(tkn.Token)
		if err != nil {
			t.Errorf("token was not saved to database - token: %v, err: %v", tkn, err)
		}
		err = DeleteAllExpiredUserTokens(u)
		if err != nil {
			t.Errorf("return error %v", err)
		}
		tkn, err = GetUserToken("expired token")
		if !errors.Is(err, utils.NotFoundError) {
			t.Errorf("user token still exist - token: %v, err: %v", tkn, err)
		}
	})

	t.Run("keep user's token not expired", func(t *testing.T) {
		u := user.NewUser("expired", "test")
		u.UUID = "2"
		s := service.NewService("test", "http://url")
		tkn := NewUserToken(u, s)
		err := tkn.Create()
		if err != nil {
			t.Errorf("faile to create token - token: %v, err: %v", tkn, err)
		}
		tkn, err = GetUserToken(tkn.Token)
		if err != nil {
			t.Errorf("token was not saved to database - token: %v, err: %v", tkn, err)
		}
		err = DeleteAllExpiredUserTokens(u)
		if err != nil {
			t.Errorf("return error %v", err)
		}
		tkn, err = GetUserToken(tkn.Token)
		if err != nil {
			t.Errorf("not expired user token was also deleted - token: %v, err: %v", tkn, err)
		}
	})
}

func TestUserToken(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()
	u := user.NewUser("test user token", "test")
	u.Create()
	s := service.NewService("test user token", "http://url")
	s.Create()
	tkn := NewUserToken(u, s)

	t.Run("Create", func(t *testing.T) {
		t.Run("create valid record", func(t *testing.T) {
			err := tkn.Create()
			if err != nil {
				t.Errorf("return error %v", err)
			}
		})

		t.Run("create existing token", func(t *testing.T) {
			err := tkn.Create()
			if !errors.Is(err, utils.DatabaseError) {
				t.Errorf("return error %v", err)
			}
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("delete existing record", func(t *testing.T) {
			err := tkn.Delete()
			if err != nil {
				t.Errorf("return error %v", err)
			}
		})

		t.Run("delete not exist token", func(t *testing.T) {
			err := tkn.Delete()
			// delete not exist record in db will not report error
			if err != nil {
				t.Errorf("return error %v", err)
			}
		})
	})

	t.Run("Expire", func(t *testing.T) {
		t.Run("success", func(t *testing.T) {
			err := tkn.Expire()
			if err != nil {
				t.Errorf("return error %v", err)
			}
			if err := tkn.Valid(); !errors.Is(err, TokenExpiredError) {
				t.Errorf("fail to expired user token")
			}
		})
	})

	t.Run("Valid", func(t *testing.T) {
		t.Run("return token expired error if token already expired", func(t *testing.T) {
			err := tkn.Valid()
			if !errors.Is(err, TokenExpiredError) {
				t.Errorf("return wrong error %v", err)
			}
		})

		t.Run("return nil if token not expired", func(t *testing.T) {
			tkn := NewUserToken(u, s)
			err := tkn.Valid()
			if err != nil {
				t.Errorf("return error %v", err)
			}
		})
	})

	t.Run("User", func(t *testing.T) {
		t.Run("existing user", func(t *testing.T) {
			tknUser, err := tkn.User()
			if err != nil || tknUser.Username != u.Username || tknUser.UUID != u.UUID {
				t.Errorf("return wrong user - user: %v, err: %v", tknUser, err)
			}
		})

		t.Run("not existing user", func(t *testing.T) {
			tkn.UserUUID = "not exist"
			tknUser, err := tkn.User()
			if !errors.Is(err, utils.NotFoundError) || tknUser.UUID != "" {
				t.Errorf("return not exist user - user: %v, err: %v", tknUser, err)
			}
		})
	})

	t.Run("Service", func(t *testing.T) {
		t.Run("existing service", func(t *testing.T) {
			tknService, err := tkn.Service()
			if err != nil || tknService.Name != s.Name || tknService.UUID != s.UUID {
				t.Errorf("return wrong service - service: %v, err: %v", tknService, err)
			}
		})

		t.Run("not existing service", func(t *testing.T) {
			tkn.ServiceUUID = "not exist"
			tknService, err := tkn.Service()
			if !errors.Is(err, utils.NotFoundError) || tknService.UUID != "" {
				t.Errorf("return not exist service - service: %v, err: %v", tknService, err)
			}
		})
	})
}
