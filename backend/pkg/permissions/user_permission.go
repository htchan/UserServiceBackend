package permissions

import (
	"errors"
	"github.com/htchan/UserService/backend/pkg/user"
)

type UserPermission struct {
	userUUID, Permission string
}

func GrantPermission(user *user.User, permission *ServicePermission) error {
	if _, err := FindUserPermissionByPermission(user, permission.Permission); err == nil {
		return errors.New("permission already grant")
	}
	userPermission := new(UserPermission)
	userPermission.userUUID = user.UUID
	userPermission.Permission = permission.Permission
	return userPermission.create()
}

func RevokePermission(permission *UserPermission) error {
	return permission.delete()
}
