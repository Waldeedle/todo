package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func AddDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/todo.db")
	if err != nil {
		fmt.Println("Could not connect to db", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Could not ping db", err)
		return nil, err
	}

	return db, nil
}
