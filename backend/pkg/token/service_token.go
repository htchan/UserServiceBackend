package token

import (
	"crypto/rand"
	"encoding/hex"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/service"
)

type ServiceToken struct {
	Token       string
	ServiceUUID string
}

var emptyServiceToken = ServiceToken{}

const (
	ServiceTokenLength = 64
)

func generateServiceToken() string {
	b := make([]byte, ServiceTokenLength)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)[:ServiceTokenLength]
}

func NewServiceToken(s service.Service) ServiceToken {
	return ServiceToken{
		ServiceUUID: s.UUID,
		Token:       generateServiceToken(),
	}
}

func GetServiceToken(tokenString string) (ServiceToken, error) {
	return queryServiceToken(
		ServiceToken{Token: tokenString}, "select service token",
		"select service_uuid, token from service_tokens where token=?",
		tokenString,
	)
}

func (t ServiceToken) Create() error {
	return utils.Execute(
		t, "create user token",
		"insert into service_tokens (service_uuid, token) values (?, ?)",
		t.ServiceUUID, t.Token,
	)
}

func (t ServiceToken) Delete() error {
	return utils.Execute(
		t, "delete user token",
		"delete from service_tokens where token=?",
		t.Token,
	)
}

func (t ServiceToken) Service() (service.Service, error) {
	return service.GetService(t.ServiceUUID)
}
