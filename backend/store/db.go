package store

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var db *sql.DB

func CloseDB() {
	db.Close()
}

func OpenDB() {
	newdb, err := sql.Open("postgres", os.Getenv("DB_CONN"))
	if err != nil {
		log.Panic(err)
	}
	err = newdb.Ping()
	if err != nil {
		log.Panic(err)
	}
	db = newdb
}
