package utils

import (
	"testing"
)

func init() {

}

func TestOpenCloseDB(t *testing.T) {
	database := GetDB()
	if database != nil {
		t.Fatalf("utils.database is not nil")
	}
	OpenDB("../../test/utils-test-data.db")
	database = GetDB()
	if database == nil {
		t.Fatalf("utils.StartDB() does not give value to users.database")
	}
	CloseDB()
	database = GetDB()
	if database != nil {
		t.Fatalf("utils.CloseDB() does not close database properly")
	}
}