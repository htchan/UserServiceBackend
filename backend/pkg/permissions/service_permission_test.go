package permissions

import (
	"testing"
	"io"
	"os"
	"github.com/htchan/UserService/backend/internal/utils"
	"github.com/htchan/UserService/backend/pkg/services"
)

func init() {
	// copy database to test environment
	source, err := os.Open("../../assets/template.db")
	utils.CheckError(err)
	destination, err := os.Create("../../test/permissions/service-permission-test-data.db")
	utils.CheckError(err)
	io.Copy(destination, source)
	source.Close()
	destination.Close()
}

func TestRegisterPermission(t *testing.T) {
	utils.OpenDB("../../test/permissions/service-permission-test-data.db")
	defer utils.CloseDB()

	service, err := services.RegisterService("reg_service", "/")
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		permission, err := RegisterPermission(service, "permission")
		if permission == nil || err != nil || permission.serviceUUID == "" {
			t.Fatalf("permissions.RegisterPermission() returns permission: %v, error: %v",
				permission, err)
		}
		resultPermission, err := FindServicePermissionByPermission(service, "permission")
		utils.CheckError(err)
		if resultPermission.Permission != permission.Permission || resultPermission.serviceUUID != permission.serviceUUID {
			t.Fatalf("permissions.RegisterPermissions() does not save permission")
		}
	})

	t.Run("fail for existing permission", func(t *testing.T) {
		permission, err := RegisterPermission(service, "permission")
		if permission != nil || err == nil {
			t.Fatalf("permissions.RegisterPermission() on existing permission returns permission: %v, error: %v",
				permission, err)
		}
	})
}

func TestingUnregisterPermission(t *testing.T) {
	utils.OpenDB("../../test/permissions/service-permission-test-data.db")
	defer utils.CloseDB()

	service, err := services.RegisterService("reg_service", "/")
	utils.CheckError(err)
	permission, err := RegisterPermission(service, "permission")
	utils.CheckError(err)

	t.Run("success", func(t *testing.T) {
		err = UnregisterPermission(service, permission)
		if err != nil {
			t.Fatalf("permissions.UnregisterPermission() returns error %v", err)
		}
	})
}
