package services

import (
	"errors"
	"github.com/google/uuid"
)

type Service struct {
	UUID string
	Name string
}

func RegisterService(name string) (*Service, error) {
	if _, err := FindServiceByName(name); err == nil {
		return nil, errors.New("service already exist")
	}
	service := new(Service)
	service.UUID = uuid.NewString()
	service.Name = name
	
	err := service.create()
	if err != nil {
		return nil, err
	}
	return service, nil
}

func UnregisterService(service *Service) error {
	return service.delete()
}