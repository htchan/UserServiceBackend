package tokens

import (
	"os"
	"io"
	"testing"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/services"
)

func init() {
	// copy database to test environment
	source, err := os.Open("../../assets/template.db")
	utils.CheckError(err)
	destination, err := os.Create("../../test/tokens/service-token-test-data.db")
	utils.CheckError(err)
	io.Copy(destination, source)
	source.Close()
	destination.Close()
}

func Test_generateServiceToken(t *testing.T) {
	utils.OpenDB("../../test/tokens/service-token-test-data.db")
	defer utils.CloseDB()
	service, err := services.RegisterService("generate_token")
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		serviceToken := generateServiceToken(service)
		if serviceToken.serviceUUID == "" || len(serviceToken.Token) != 64 {
			t.Fatalf("tokens.generateServiceToken() returns serviec token %v", serviceToken)
		}
	})
}

func TestLoadServiceToken(t *testing.T) {
	utils.OpenDB("../../test/tokens/service-token-test-data.db")
	defer utils.CloseDB()

	t.Run("success for existing service token", func(t *testing.T) {
		service, err := services.RegisterService("owner_token")
		utils.CheckError(err)
		serviceToken := generateServiceToken(service)
		err = serviceToken.create()
		utils.CheckError(err)
		actualServiceToken, err := LoadServiceToken(service)
		if actualServiceToken == nil || err != nil {
			t.Fatalf("tokens.LoadServiceToken() returns service token %v, err %v",
				actualServiceToken, err)
		}
		if actualServiceToken.serviceUUID == "" || actualServiceToken.Token != serviceToken.Token {
			t.Fatalf("actual token %v\ndifferent from\nexpect token %v",
				actualServiceToken, serviceToken)
		}
	})

	t.Run("success for not exist service token", func(t *testing.T) {
		service, err := services.RegisterService("load_token")
		utils.CheckError(err)
		actualServiceToken, err := LoadServiceToken(service)
		if actualServiceToken == nil || err != nil ||
			actualServiceToken.serviceUUID == "" ||
			len(actualServiceToken.Token) != 64 {
			t.Fatalf("tokens.LoadServiceToken() returns service token %v, err %v",
				actualServiceToken, err)
		}
	})
}

func TestDeleteServiceTokens(t *testing.T) {
	utils.OpenDB("../../test/tokens/service-token-test-data.db")
	defer utils.CloseDB()

	t.Run("success", func(t *testing.T) {
		service, err := services.RegisterService("delete_token")
		utils.CheckError(err)
		serviceToken, err := LoadServiceToken(service)
		utils.CheckError(err)
		err = DeleteServiceTokens(service)
		if err != nil {
			t.Fatalf("tokens.DeleteServiceTokens() return err %v", err)
		}
		newServiceToken, err := LoadServiceToken(service)
		utils.CheckError(err)
		if newServiceToken.Token == serviceToken.Token {
			t.Fatalf("tokens.DeleteServiceTokens() does not remove token in database")
		}
	})
}

func TestRenewServiceToken(t *testing.T) {
	utils.OpenDB("../../test/tokens/service-token-test-data.db")
	defer utils.CloseDB()
	
	t.Run("success", func(t *testing.T) {
		service, err := services.RegisterService("renew_token")
		utils.CheckError(err)
		serviceToken, err := LoadServiceToken(service)
		utils.CheckError(err)
		err = RenewServiceToken(service)
		if err != nil {
			t.Fatalf("tokens.RenewServiceToken() return err %v", err)
		}
		newServiceToken, err := LoadServiceToken(service)
		utils.CheckError(err)
		if newServiceToken.Token == serviceToken.Token {
			t.Fatalf("tokens.DeleteServiceTokens() does not remove token in database")
		}
	})
}