package tokens

import (
	"errors"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/services"
	"github.com/htchan/UserService/backend/pkg/users"
)

func (userToken UserToken) create() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("insert into user_tokens (user_uuid, token, created_date, duration) values (?, ?, ?, ?)",
		userToken.userUUID, userToken.Token, userToken.generateDate,
		userToken.duration)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (userToken UserToken) delete() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from user_tokens where user_uuid=? and token=?",
		userToken.userUUID, userToken.Token)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func FindUserTokenByTokenStr(tokenStr string) (*UserToken, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select user_uuid, created_date, duration from user_tokens where token=?",
		tokenStr)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		userToken := new(UserToken)
		userToken.Token = tokenStr
		rows.Scan(&userToken.userUUID, &userToken.generateDate, &userToken.duration)
		return userToken, nil
	}
	return nil, errors.New("invalid token")
}

func FindUserTokenByUser(user *users.User) (*UserToken, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select token, created_date, duration from user_tokens where user_uuid=?",
		user.UUID)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		userToken := new(UserToken)
		userToken.userUUID = user.UUID
		rows.Scan(&userToken.Token, &userToken.generateDate, &userToken.duration)
		return userToken, nil
	}
	return nil, errors.New("invalid user")
}

func FindUserByTokenStr(tokenStr string) (*users.User, error) {
	token, err := FindUserTokenByTokenStr(tokenStr)
	if err != nil {
		return nil, err
	}
	return users.FindUserByUUID(token.userUUID)
}

func (serviceToken ServiceToken) create() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil
	}
	_, err = tx.Exec("insert into service_tokens (service_uuid, token) values (?, ?)",
		serviceToken.serviceUUID, serviceToken.Token)
	if err != nil {
		tx.Rollback()
		return nil
	}
	return tx.Commit()
}

func (serviceToken ServiceToken) delete() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		tx.Rollback()
		return nil
	}
	_, err = tx.Exec("delete from service_tokens where token=?", serviceToken.Token)
	if err != nil {
		tx.Rollback()
		return nil
	}
	return tx.Commit()
}

func (serviceToken ServiceToken) update() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil
	}
	_, err = tx.Exec("update service_tokens set token=? where service_uuid=?",
		serviceToken.Token, serviceToken.serviceUUID)
	if err != nil {
		tx.Rollback()
		return nil
	}
	return tx.Commit()
}

func FindServiceTokenByTokenStr(tokenStr string) (*ServiceToken, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select service_uuid from service_tokens where token=?",
		tokenStr)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		serviceToken := new(ServiceToken)
		serviceToken.Token = tokenStr
		rows.Scan(&serviceToken.serviceUUID)
		return serviceToken, nil
	}
	return nil, errors.New("invalid token")
}

func FindServiceTokenByService(service *services.Service) (*ServiceToken, error) {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.Query("select token from service_tokens where service_uuid=?",
		service.UUID)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		serviceToken := new(ServiceToken)
		serviceToken.serviceUUID = service.UUID
		rows.Scan(&serviceToken.Token)
		return serviceToken, nil
	}
	return nil, errors.New("invalid service")
}

func FindServiceByTokenStr(tokenStr string) (*services.Service, error) {
	token, err := FindServiceTokenByTokenStr(tokenStr)
	if err != nil {
		return nil, err
	}
	return services.FindServiceByUUID(token.serviceUUID)
}