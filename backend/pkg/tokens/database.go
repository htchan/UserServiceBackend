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
	_, err = tx.Exec("insert into user_tokens (username, token, created_date, duration) values (?, ?, ?, ?)",
		userToken.Username, userToken.Token, userToken.generateDate, userToken.duration)
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
	_, err = tx.Exec("delete from user_tokens where username=? and token=?",
		userToken.Username, userToken.Token)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func FindUserTokenByTokenStr(tokenStr string) (*UserToken, error) {
	tx, err := utils.GetDB().Begin()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("select username, created_date, duration from user_tokens where token=?",
		tokenStr)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		userToken := new(UserToken)
		userToken.Token = tokenStr
		rows.Scan(&userToken.Username, &userToken.generateDate, &userToken.duration)
		return userToken, nil
	}
	return nil, errors.New("invalid token")
}

func FindUserTokenByUser(user users.User) (*UserToken, error) {
	tx, err := utils.GetDB().Begin()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("select token, created_date, duration from user_tokens where username=?",
		user.Username)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		userToken := new(UserToken)
		userToken.Username = user.Username
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
	return users.FindUserByName(token.Username)
}

func (serviceToken ServiceToken) create() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil
	}
	_, err = tx.Exec("insert into service_tokens (service_name, token) values (?, ?)",
		serviceToken.serviceName, serviceToken.Token)
	if err != nil {
		return nil
	}
	return tx.Commit()
}

func (serviceToken ServiceToken) delete() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil
	}
	_, err = tx.Exec("delete from service_tokens where token=?", serviceToken.Token)
	if err != nil {
		return nil
	}
	return tx.Commit()
}

func (serviceToken ServiceToken) update() error {
	tx, err := utils.GetDB().Begin()
	if err != nil {
		return nil
	}
	_, err = tx.Exec("update service_tokens set token=? where service_name=?",
		serviceToken.Token, serviceToken.serviceName)
	if err != nil {
		return nil
	}
	return tx.Commit()
}

func FindServiceTokenByTokenStr(tokenStr string) (*ServiceToken, error) {
	tx, err := utils.GetDB().Begin()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("select service_name from service_tokens where token=?",
		tokenStr)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		serviceToken := new(ServiceToken)
		serviceToken.Token = tokenStr
		rows.Scan(&serviceToken.serviceName)
		return serviceToken, nil
	}
	return nil, errors.New("invalid token")
}

func FindServiceTokenByService(service services.Service) (*ServiceToken, error) {
	tx, err := utils.GetDB().Begin()
	defer tx.Rollback()
	if err != nil {
		return nil, err
	}
	rows, err := tx.Query("select token from service_tokens where service_name=?",
		service.Name)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		serviceToken := new(ServiceToken)
		serviceToken.serviceName = service.Name
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
	return services.FindServiceByName(token.serviceName)
}