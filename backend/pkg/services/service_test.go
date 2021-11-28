package services

import (
	"testing"
	"os"
	"io"
	"github.com/htchan/UserService/backend/internal/utils"
)

func init() {
	// copy database to test environment
	source, err := os.Open("../../assets/template.db")
	utils.CheckError(err)
	destination, err := os.Create("../../test/services/services-test-data.db")
	utils.CheckError(err)
	io.Copy(destination, source)
	source.Close()
	destination.Close()
}

func TestRegisterService(t *testing.T) {
	utils.OpenDB("../../test/services/services-test-data.db")
	defer utils.CloseDB()
	t.Run("success", func(t *testing.T) {
		service, err := RegisterService("reg_service", "some_url")
		if service == nil || err != nil ||
			service.Name != "reg_service" || service.UUID == "" ||
			service.Url != "some_url" {
			t.Fatalf("services.RegisterService return service %v, error %v",
				service, err)
		}
		if _, err := FindServiceByName("reg_service"); err != nil {
			t.Fatalf("service.RegisterService cannot save service")
		}
	})

	t.Run("existing service", func(t *testing.T) {
		service, err := RegisterService("reg_service", "some_url")
		if service != nil || err == nil {
			t.Fatalf("services.RegisterService return service %v, error %v",
				service, err)
		}
	})
}

func TestUnregisterService(t *testing.T) {
	utils.OpenDB("../../test/services/services-test-data.db")
	defer utils.CloseDB()
	service, err := RegisterService("unreg_service", "some_url")
	utils.CheckError(err)

	t.Run("exist service", func(t *testing.T) {
		err = UnregisterService(service)
		if err != nil {
			t.Fatalf("services.UnregisterService on exist service return error : %v", err)
		}
		if _, err := FindServiceByName("unreg_service"); err == nil {
			t.Fatalf("services.UnregisterService does not delete service")
		}
	})

	t.Run("not exist service will not return error", func(t *testing.T) {
		err = UnregisterService(service)
		if err != nil {
			t.Fatalf("services.UnregisterService on not exist service return error : %v", err)
		}
	})
}