package permissions

// import (
// 	"testing"
// 	"io"
// 	"os"
// 	"github.com/htchan/UserService/backend/internal/utils"
// 	"github.com/htchan/UserService/backend/pkg/user"
// 	"github.com/htchan/UserService/backend/pkg/service"
// )

// func init() {
// 	// copy database to test environment
// 	source, err := os.Open("../../assets/template.db")
// 	utils.CheckError(err)
// 	destination, err := os.Create("../../test/permissions/user-permission-test-data.db")
// 	utils.CheckError(err)
// 	io.Copy(destination, source)
// 	source.Close()
// 	destination.Close()
// }

// func TestGrantPermission(t *testing.T) {
// 	utils.OpenDB("../../test/permissions/user-permission-test-data.db")
// 	defer utils.CloseDB()

// 	user, err := user.Signup("grant_user", "password")
// 	utils.CheckError(err)
// 	service, err := service.RegisterService("grant_service", "/")
// 	utils.CheckError(err)
// 	permission, err := RegisterPermission(service, "grant_permission")
// 	utils.CheckError(err)

// 	t.Run("success", func(t *testing.T) {
// 		err := GrantPermission(user, permission)
// 		if err != nil {
// 			t.Fatalf("permissions.GrantPermission() returns err %v", err)
// 		}
// 		resultPermissions, err := FindUserPermissionsByUser(user)
// 		utils.CheckError(err)
// 		if len(resultPermissions) != 1 && resultPermissions[0].Permission != "grant_permission" {
// 			t.Fatalf("permissions.GrantPermission() save permission as %v",
// 				resultPermissions)
// 		}
// 	})

// 	t.Run("already exist", func(t *testing.T) {
// 		err := GrantPermission(user, permission)
// 		if err == nil {
// 			t.Fatalf("permissions.GrantPermission() does not return error on granted permissions")
// 		}
// 	})
// }

// func TestRevokePermission(t *testing.T) {
// 	utils.OpenDB("../../test/permissions/user-permission-test-data.db")
// 	defer utils.CloseDB()

// 	user, err := user.Signup("revoke_user", "password")
// 	utils.CheckError(err)
// 	service, err := service.RegisterService("revoke_service", "/")
// 	utils.CheckError(err)
// 	servicePermission, err := RegisterPermission(service, "revoke_permission")
// 	utils.CheckError(err)
// 	err = GrantPermission(user, servicePermission)
// 	utils.CheckError(err)
// 	userPermission, err := FindUserPermissionByPermission(user, servicePermission.Permission)
// 	utils.CheckError(err)

// 	t.Run("success", func(t *testing.T) {
// 		err = RevokePermission(userPermission)
// 		if err != nil {
// 			t.Fatalf("permissions.RevokePermission() returns err %v", err)
// 		}
// 		resultPermissions, err := FindUserPermissionsByUser(user)
// 		utils.CheckError(err)
// 		if len(resultPermissions) != 0 {
// 			t.Fatalf("permissions.RevokePermission() save permission as %v",
// 				resultPermissions)
// 		}
// 	})
// }
