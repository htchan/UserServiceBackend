package utils

import (
	"log"
	"errors"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func OpenDB(dbLocation string) {
	var err error
	database, err = sql.Open("sqlite3", dbLocation)
    CheckError(err)
}

func CloseDB() {
	err := database.Close()
	CheckError(err)
	database = nil
}

func GetDB() *sql.DB {
	return database
}

var DatabaseError = errors.New("database error")

func NewDatabaseError(operation string, data interface{}, err error) error {
	log.Printf("[%v] database error - data: %v, err: %v", operation, data, err)
	return DatabaseError
}

func Execute(data interface{}, operation, command string, params ...interface{}) error {
	tx, err := GetDB().Begin()
	if err != nil {
		return NewDatabaseError(operation, data, err)
	}

	commited := false
	defer func() {
		if !commited {
			tx.Rollback()
		}
	}()

	_, err = tx.Exec(command, params...)
	if err != nil {
		return NewDatabaseError(operation, data, err)
	}

	err = tx.Commit()
	if err != nil {
		return NewDatabaseError(operation, data, err)
	}
	commited = true
	return nil
}

var NotFoundError = errors.New("not found error")

func NewNotFoundError(operation string, data interface{}, err error) error {
	log.Printf("[%v] record not found - data: %v, err: %v", operation, data, err)
	return NotFoundError
}

var InvalidRecordError = errors.New("invalid record")

func NewInvalidRecordError(content string) error {
	log.Printf("invalid record: %v", content)
	return InvalidRecordError
}