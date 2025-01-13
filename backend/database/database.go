package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE NOT NULL,
		password_hash TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	createArticlesTable := `
	CREATE TABLE IF NOT EXISTS articles (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		content TEXT NOT NULL,
		user_id INTEGER NOT NULL,
		publish_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES users(id)
	);`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(createArticlesTable)
	if err != nil {
		log.Fatal(err)
	}
}
