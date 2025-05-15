package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
}

func GetDB() *sql.DB {
	return db
}

// Additional CRUD operations can be added here as needed.