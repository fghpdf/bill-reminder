package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./sqlite/withdrawals.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS withdrawals (
		id TEXT PRIMARY KEY,
		date TEXT,
		description TEXT,
		amount INTEGER,
		created_at TEXT
	);`
	_, err = DB.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
}
