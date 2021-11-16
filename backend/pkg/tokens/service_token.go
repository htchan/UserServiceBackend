package tokens

import (
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/services"
)

type ServiceToken struct {
	serviceUUID string
	Token string
}

func generateServiceToken(service services.Service) *ServiceToken {
	serviceToken := new(ServiceToken)
	for true {
		serviceToken.Token = utils.RandomString(64)
		if _, err := FindUserTokenByTokenStr(serviceToken.Token); err != nil {
			break
		}
	}
	serviceToken.serviceUUID = service.UUID
	return serviceToken
}

func LoadServiceToken(service services.Service) (*ServiceToken, error) {
	token, err := FindServiceTokenByService(service)
	if err == nil {
		return token, nil
	}
	serviceToken := generateServiceToken(service)
	err = serviceToken.create()
	if err != nil {
		return nil, err
	}
	return serviceToken, nil
}

func DeleteServiceTokens(service services.Service) error {

	tx, err := utils.GetDB().Begin()
	if err != nil {
		return err
	}
	_, err = tx.Exec("delete from service_tokens where service_uuid=?", service.UUID)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func RenewServiceToken(service services.Service) error {
	token, err := FindServiceTokenByService(service)
	if err != nil {
		return err
	}
	token.Token = utils.RandomString(64)
	return token.update()
}