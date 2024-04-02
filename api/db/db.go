package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var Database *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite3", "api.db")
	Database = db

	if err != nil {
		panic("Could not connect to database.")
	}

	Database.SetMaxOpenConns(10)
	Database.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`

	_, err := Database.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table.")
	}

	createSuggestionsTable := `
	CREATE TABLE IF NOT EXISTS suggestions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT NOT NULL,
		text TEXT NOT NULL,
		FOREIGN KEY(username) REFERENCES users(username)
	)
	`

	_, err = Database.Exec(createSuggestionsTable)

	if err != nil {
		panic("Could not create suggestions table.")
	}
}
