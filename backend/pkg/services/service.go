package services

import (
	"errors"
	"strings"
	"github.com/google/uuid"
)

type Service struct {
	UUID string
	Name string
	Url string
}

func RegisterService(name, url string) (*Service, error) {
	if _, err := FindServiceByName(name); err == nil {
		return nil, errors.New("service already exist")
	}
	if !strings.HasSuffix(url, "/") {
		return nil, errors.New("url have end with /")
	}
	service := new(Service)
	service.UUID = uuid.NewString()
	service.Name = name
	service.Url = url
	
	err := service.create()
	if err != nil {
		return nil, err
	}
	return service, nil
}

func UnregisterService(service *Service) error {
	return service.delete()
}