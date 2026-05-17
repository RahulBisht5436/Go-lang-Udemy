package db

import (
	// Go uses this SQL package
	"database/sql"
	// Pure-Go SQLite driver (no CGO / no C compiler required)
	_ "modernc.org/sqlite"
)

// Global DB variable , type sql.DB pointer type
var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Error in the DB initialization: " + err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic("DB unreachable: " + err.Error())
	}

	// Defines the max number of connection inside the DB
	DB.SetMaxOpenConns(10)

	//Defines then idle connection number
	DB.SetMaxIdleConns(5)

	// initiate the DB
	createTables()

}
func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime TEXT NOT NULL,
	user_id INTEGER

	) 
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Not able to create the table: " + err.Error())
	}
}
