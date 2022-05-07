package permissions

import (
	"errors"
	"github.com/htchan/UserService/backend/pkg/service"
)

type ServicePermission struct {
	serviceUUID, Permission string
}

func RegisterPermission(service *service.Service, permissionStr string) (*ServicePermission, error) {
	if _, err := FindServicePermissionByPermission(service, permissionStr); err == nil {
		return nil, errors.New("permission already registered")
	}
	permission := new(ServicePermission)
	permission.serviceUUID = service.UUID
	permission.Permission = permissionStr
	err := permission.create()
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func UnregisterPermission(service *service.Service, permission *ServicePermission) error {
	if _, ok := FindServicePermissionByPermission(service, permission.Permission); ok != nil {
		return errors.New("permission not found")
	}
	return permission.delete()
}