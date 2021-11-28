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
	_, err = tx.Exec("insert into services (uuid, name, url) values(?, ?, ?)",
		service.UUID, service.Name, service.Url)
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
	_, err = tx.Exec("delete from services where uuid=?", service.UUID)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func FindServiceByName(serviceName string) (*Service, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select uuid, name, url from services where name=?", serviceName)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		service := new(Service)
		rows.Scan(&service.UUID, &service.Name, &service.Url)
		return service, nil
	}
	return nil, errors.New("service not found")
}

func FindServiceByUUID(uuid string) (*Service, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select uuid, name, url from services where uuid=?",
		uuid)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		service := new(Service)
		rows.Scan(&service.UUID, &service.Name, &service.Url)
		return service, nil
	}
	return nil, errors.New("service not found")
}