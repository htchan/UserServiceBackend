package service

import (
	"strings"
	"github.com/google/uuid"
)

type Service struct {
	UUID string
	Name string
	Url string
}

var emptyService = Service{}

func RegisterService(name, url string) (Service, error) {
	if _, err := FindServiceByName(name); err == nil {
		return emptyService, ServiceAlreadyExistError(name)
	}
	if !strings.HasSuffix(url, "/") {
		return emptyService, InvalidUrlError("url have end with /")
	}
	service := Service{
		UUID: uuid.NewString(),
		Name: name,
		Url: url,
	}

	err := service.create()
	if err != nil {
		return emptyService, err
	}
	return service, nil
}

func UnregisterService(service Service) error {
	return service.delete()
}