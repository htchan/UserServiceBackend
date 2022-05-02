package service

import (
	"testing"
	"os"
	"io"
	"errors"
	"github.com/htchan/UserService/backend/internal/utils"
)

func init() {
	// copy database to test environment
	source, err := os.Open("../../assets/test_data.db")
	utils.CheckError(err)
	destination, err := os.Create("./services-test-data.db")
	utils.CheckError(err)
	io.Copy(destination, source)
	source.Close()
	destination.Close()
}

func TestNewService(t *testing.T) {
	t.Run("generate service instance", func (t *testing.T) {
		s := NewService("test name", "http://url")
		if len(s.UUID) == 0 || s.Name != "test name" || s.URL != "http://url" {
			t.Errorf("generate wrong instance: %v", s)
		}
	})
}

func TestGetService(t *testing.T) {
	utils.OpenDB("./services-test-data.db")
	defer utils.CloseDB()

	t.Run("query existing service", func (t *testing.T) {
		s, err := GetService("1")
		if err != nil || s.UUID != "1" || s.Name != "user service" || s.URL != "http://url" {
			t.Errorf("failed to get existing service - service: %v, err: %v", s, err)
		}
	})

	t.Run("query missing service", func (t *testing.T) {
		s, err := GetService("not exist")
		if err == nil || !errors.Is(err, utils.NotFoundError) || s != emptyService {
			t.Errorf("not reporting error - service: %v, err: %v", s, err)
		}
	})
}

func TestService(t *testing.T) {
	utils.OpenDB("./services-test-data.db")
	defer utils.CloseDB()

	s := NewService("test name", "http://url")

	t.Run("Create", func (t *testing.T) {
		t.Run("create valid service", func (t *testing.T) {
			err := s.Create()
			if err != nil {
				t.Errorf("create return error: %v", err)
			}
		})

		t.Run("create existing service", func (t *testing.T) {
			err := s.Create()
			if err == nil || !errors.Is(err, utils.DatabaseError) {
				t.Errorf("create exist service not return expected error: %v", err)
			}
		})
	})

	t.Run("Delete", func (t *testing.T) {
		t.Run("delete existing service", func (t *testing.T) {
			s.Create()
			err := s.Delete()
			if err != nil {
				t.Errorf("create return error: %v", err)
			}
		})

		t.Run("delete not exist service", func (t *testing.T) {
			err := s.Delete()
			// delete not exist record will not return error
			if err != nil {
				t.Errorf("create return error: %v", err)
			}
		})
	})

	t.Run("Valid", func (t *testing.T) {
		t.Run("service url start with http", func (t *testing.T) {
			s := NewService("valid service", "http://url")
			if err := s.Valid(); err != nil {
				t.Errorf("valid return error: %v", err)
			}
		})
		
		t.Run("service url start with ftp", func (t *testing.T) {
			s := NewService("invalid service", "ftp://url")
			if err := s.Valid(); err == nil || !errors.Is(err, utils.InvalidRecordError) {
				t.Errorf("valid return error: %v", err)
			}
		})
	})
}
