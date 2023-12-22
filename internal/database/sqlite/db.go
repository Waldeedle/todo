package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func AddDB() {
	db, err := sql.Open(
		"sqlite3",
		"user:password@tcp(127.0.0.1:3306)/hello",
	)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
