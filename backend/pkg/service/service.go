package service

import (
	"os"
	"fmt"
	"strings"
	"github.com/google/uuid"
	"github.com/htchan/UserService/backend/internal/utils"
)

type Service struct {
	UUID string
	Name string
	URL string
}

var emptyService = Service{}

var (
	defaultUserServiceName = "user_service"
	defaultUserServiceURL = os.Getenv("DEFAULT_USER_SERVICE_URL")
	defaultUserService = emptyService
)

func NewService(name, url string) Service {
	return Service{
		UUID: uuid.NewString(),
		Name: name,
		URL: url,
	}
}

func GetService(uuid string) (Service, error) {
	return queryService(
		Service{UUID: uuid}, "query service by uuid",
		"select uuid, name, url from services where uuid=?",
		uuid,
	)
}

func GetServiceByName(name string) (Service, error) {
	return queryService(
		Service{Name: name}, "query service by name",
		"select uuid, name, url from services where name=?",
		name,
	)
}

func DefaultUserService() Service {
	if defaultUserService != emptyService {
		return defaultUserService
	}
	var err error
	defaultUserService, err = GetServiceByName(defaultUserServiceName)
	if err != nil {
		defaultUserService = NewService(defaultUserServiceName, defaultUserServiceURL)
		defaultUserService.Create()
	}
	return defaultUserService
}

func (s Service) Create() error {
	return utils.Execute(
		s, "create service",
		"insert into services (uuid, name, url) values (?, ?, ?)",
		s.UUID, s.Name, s.URL,
	)
}

func (s Service) Delete() error {
	return utils.Execute(
		s, "delete service",
		"delete from services where uuid=?",
		s.UUID,
	)
}

func (s Service) Valid() error {
	if !strings.HasPrefix(s.URL, "http") {
		return utils.NewInvalidRecordError("url have end with /")
	}
	return nil
}

func (s Service) RedirectURL(token string) string {
	return fmt.Sprintf("%v?token=%v", s.URL, token)
}
