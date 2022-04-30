package service

import (
	"github.com/htchan/UserService/backend/internal/utils"
	"fmt"
)

var USER_SERVICE = emptyService

func execute(operation, model string, sql string, params ...interface{}) error {
	tx, err := utils.GetDB().Begin()
	if err != nil { return DatabaseError{operation, model, err} }
	commited := false
	defer func() { if !commited { tx.Rollback() } }()

	_, err = tx.Exec(sql, params...)
	if err != nil {
		return DatabaseError{operation, model, err}
	}
	err = tx.Commit()
	if err != nil { return DatabaseError{operation, model, err} }
	commited = true
	return nil
}

func (service Service) create() error {
	return execute("create", "service", "insert into services (uuid, name, url) values(?, ?, ?)",
		service.UUID, service.Name, service.Url)
}

func (service Service) delete() error {
	return execute("delete", "service", "delete from services where uuid=?", service.UUID)
}

func FindServiceByName(serviceName string) (Service, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return emptyService, DatabaseError{"select", "service", err}
	}
	defer tx.Rollback()
	rows, err := tx.Query("select uuid, url from services where name=?", serviceName)
	if err != nil {
		return emptyService, DatabaseError{"select", "service", err}
	}
	if rows.Next() {
		service := Service{ Name: serviceName }
		rows.Scan(&service.UUID, &service.Url)
		return service, nil
	}
	return emptyService, ServiceNotFoundError(serviceName)
}

func FindServiceByUUID(uuid string) (Service, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return emptyService, DatabaseError{"select", "service", err}
	}
	defer tx.Rollback()
	rows, err := tx.Query("select name, url from services where uuid=?",
		uuid)
	if err != nil {
		return emptyService, DatabaseError{"select", "service", err}
	}
	if rows.Next() {
		service := Service{ UUID: uuid }
		rows.Scan(&service.Name, &service.Url)
		return service, nil
	}
	return emptyService, ServiceNotFoundError(uuid)
}

func UserService() Service {
	if USER_SERVICE != emptyService { return USER_SERVICE }
	USER_SERVICE, err := FindServiceByName("UserService")
	if err != nil {
		USER_SERVICE, err = RegisterService("UserService", "/")
		if err != nil { fmt.Println(err) }
	}
	return USER_SERVICE
}