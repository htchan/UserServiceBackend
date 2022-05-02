package token

import (
	"errors"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/service"
	"io"
	"os"
	"testing"
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

func Test_generateServiceToken(t *testing.T) {
	t.Run("generate string with length 64", func(t *testing.T) {
		s := generateServiceToken()
		if len(s) != ServiceTokenLength {
			t.Errorf("token length is %v", len(s))
		}
	})
}

func TestNewServiceToken(t *testing.T) {
	s := service.NewService("test", "http://url")
	t.Run("generate service token instance", func(t *testing.T) {
		tkn := NewServiceToken(s)
		if tkn.ServiceUUID != s.UUID ||
			len(tkn.Token) != ServiceTokenLength {
			t.Errorf("wront token instance: %v", tkn)
		}
	})
}

func TestGetServiceToken(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()

	t.Run("get existing service token", func(t *testing.T) {
		tkn, err := GetServiceToken("token")
		if err != nil || tkn.ServiceUUID != "1" ||
			tkn.Token != "token" {
			t.Errorf("query wrong token - token: %v, err: %v", tkn, err)
		}
	})

	t.Run("get not exist service token", func(t *testing.T) {
		tkn, err := GetServiceToken("not exist")
		if !errors.Is(err, utils.NotFoundError) {
			t.Errorf("query wrong token - token: %v, err: %v", tkn, err)
		}
	})
}

func TestServiceToken(t *testing.T) {
	utils.OpenDB("./token-test-data.db")
	defer utils.CloseDB()
	s := service.NewService("test service token", "http://url")
	s.Create()
	tkn := NewServiceToken(s)

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
