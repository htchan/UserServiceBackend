package permissions

import (
	"errors"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/user"
	"github.com/htchan/UserService/backend/pkg/service"
)

func (userPermission UserPermission) create() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into user_permissions (user_uuid, permission) values (?, ?)",
		userPermission.userUUID, userPermission.Permission)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (userPermission UserPermission) delete() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from user_permissions where user_uuid=? and permission=?",
		userPermission.userUUID, userPermission.Permission)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func FindUserPermissionsByUser(user *user.User) ([]*UserPermission, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select user_uuid, permission from user_permissions where user_uuid=?",
		user.UUID)
	if err != nil {
		return nil, err
	}
	permissions := make([]*UserPermission, 0)
	for rows.Next() {
		userPermission := new(UserPermission)
		rows.Scan(&userPermission.userUUID, &userPermission.Permission)
		permissions = append(permissions, userPermission)
	}
	return permissions, nil
}

func FindUserPermissionByPermission(user *user.User, permissionStr string) (*UserPermission, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select user_uuid, permission from user_permissions where user_uuid=? and permission=?",
		user.UUID, permissionStr)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		userPermission := new(UserPermission)
		rows.Scan(&userPermission.userUUID, &userPermission.Permission)
		return userPermission, nil
	}
	return nil, errors.New("permission not found")
}

func (servicePermission ServicePermission) create() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into service_permissions (service_uuid, permission) values (?, ?)",
		servicePermission.serviceUUID, servicePermission.Permission)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (servicePermission ServicePermission) delete() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from service_permissions where permission=?",
		servicePermission.Permission)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = tx.Exec("delete from user_permissions where permission=?",
		servicePermission.Permission)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()

}

func FindServicePermissionsByService(service *service.Service) ([]ServicePermission, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select service_uuid, permission from service_permissions where service_uuid=?",
		service.UUID)
	if err != nil {
		return nil, err
	}
	permissions := make([]ServicePermission, 0)
	for rows.Next() {
		servicePermission := new(ServicePermission)
		rows.Scan(&servicePermission.serviceUUID, &servicePermission.Permission)
		permissions = append(permissions, *servicePermission)
	}
	return permissions, nil
}

func FindServicePermissionByPermission(service *service.Service, permissionStr string) (*ServicePermission, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select service_uuid, permission from service_permissions where permission=? and service_uuid=?",
		permissionStr, service.UUID)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		servicePermission := new(ServicePermission)
		rows.Scan(&servicePermission.serviceUUID, &servicePermission.Permission)
		return servicePermission, nil
	}
	return nil, errors.New("permission not found")
}
