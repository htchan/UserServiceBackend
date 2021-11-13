package permissions

import (
	"errors"
	"github.com/htchan/UserService/backend/pkg/users"
)

type UserPermission struct {
	username, Permission string
}

func GrantPermission(user users.User, permission ServicePermission) error {
	if _, err := FindUserPermissionByPermission(user, permission.Permission); err == nil {
		return errors.New("permission already grant")
	}
	userPermission := new(UserPermission)
	userPermission.username = user.Username
	userPermission.Permission = permission.Permission
	return userPermission.create()
}

func RevokePermission(permission UserPermission) error {
	return permission.delete()
}
