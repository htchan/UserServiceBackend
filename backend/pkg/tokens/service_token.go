package tokens

import (
	"github.com/htchan/UserService/internal/utils"
	"github.com/htchan/UserService/pkg/services"
)

type ServiceToken struct {
	serviceName string
	Token string
}

func generateServiceToken(service services.Service) *ServiceToken {
	serviceToken := new(ServiceToken)
	serviceToken.Token = utils.RandomString(64)
	serviceToken.serviceName = service.Name
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
	_, err = tx.Exec("delete from service_tokens where service_name=?", service.Name)
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