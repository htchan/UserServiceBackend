package permissions

import (
	"errors"
	"github.com/htchan/UserService/backend/pkg/services"
)

type ServicePermission struct {
	serviceUUID, Permission string
}

func RegisterPermission(service *services.Service, permissionStr string) (*ServicePermission, error) {
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

func UnregisterPermission(service *services.Service, permission *ServicePermission) error {
	if _, ok := FindServicePermissionByPermission(service, permission.Permission); ok != nil {
		return errors.New("permission not found")
	}
	return permission.delete()
}