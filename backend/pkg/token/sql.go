package token

import (
	"github.com/htchan/UserService/backend/internal/utils"
)

func queryUserToken(t UserToken, operation, query string, params ...interface{}) (UserToken, error) {
	rows, err := utils.GetDB().Query(query, params...)
	if err != nil {
		return emptyUserToken, utils.NewDatabaseError(operation, t, err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&t.UserUUID, &t.ServiceUUID, &t.Token, &t.generateDate, &t.duration)
		return t, nil
	}
	return emptyUserToken, utils.NewNotFoundError(operation, t, err)
}

func queryUserTokens(t UserToken, operation, query string, params ...interface{}) chan UserToken {
	userTokenChan := make(chan UserToken)
	go func() {
		defer close(userTokenChan)
		rows, err := utils.GetDB().Query(query, params...)
		if err != nil {
			return
		}
		defer rows.Close()
		for rows.Next() {
			var t UserToken
			rows.Scan(&t.UserUUID, &t.ServiceUUID, &t.Token, &t.generateDate, &t.duration)
			userTokenChan <- t
		}
	}()
	return userTokenChan
}

func queryServiceToken(t ServiceToken, operation, query string, params ...interface{}) (ServiceToken, error) {
	rows, err := utils.GetDB().Query(query, params...)
	if err != nil {
		return emptyServiceToken, utils.NewDatabaseError(operation, t, err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&t.ServiceUUID, &t.Token)
		return t, nil
	}
	return emptyServiceToken, utils.NewNotFoundError(operation, t, err)
}
