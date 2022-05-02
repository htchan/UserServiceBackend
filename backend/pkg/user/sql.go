package user

import (
	"github.com/htchan/UserService/backend/internal/utils"
)

func queryUser(u User, operation, query string, params ...interface{}) (User, error) {
	rows, err := utils.GetDB().Query(query, params...)
	if err != nil {
		return emptyUser, utils.NewDatabaseError(operation, u, err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&u.UUID, &u.Username, &u.EncryptedPassword)
		return u, nil
	}
	return emptyUser, utils.NewNotFoundError(operation, u, err)
}
