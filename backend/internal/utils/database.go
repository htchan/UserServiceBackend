package utils

import (
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
