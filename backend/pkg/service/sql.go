package service

import (
	"github.com/htchan/UserService/backend/internal/utils"
)

func queryService(s Service, operation, query string, params ...interface{}) (Service, error) {
	rows, err := utils.GetDB().Query(query, params...)
	if err != nil {
		return emptyService, utils.NewDatabaseError(operation, s, err)
	}
	defer rows.Close()
	if rows.Next() {
		rows.Scan(&s.UUID, &s.Name, &s.URL)
		return s, nil
	}
	return emptyService, utils.NewNotFoundError(operation, s, err)
}
