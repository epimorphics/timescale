package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

var db *sql.DB

func CloseDB() {
	db.Close()
}

func OpenDB(DatabaseConnection string) {
	newdb, err := sql.Open("postgres", DatabaseConnection)
	if err != nil {
		log.Panic(err)
	}
	err = newdb.Ping()
	if err != nil {
		log.Panic(err)
	}
	db = newdb
}
