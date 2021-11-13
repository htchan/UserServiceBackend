package services

import (
	"errors"
	"github.com/htchan/UserService/backend/internal/utils"
)

func (service Service) create() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into services (name) values(?)", service.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (service Service) delete() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from services where name=?", service.Name)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func FindServiceByName(serviceName string) (*Service, error) {
	tx, err := utils.GetDB().Begin()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("select name from services where name=?", serviceName)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		service := new(Service)
		rows.Scan(&service.Name)
		return service, nil
	}
	return nil, errors.New("service not found")
}