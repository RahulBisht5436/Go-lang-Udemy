package db

import (
	// Go uses this SQL package
	"database/sql"
	"errors"
	"fmt"
	"time"

	// Pure-Go SQLite driver (no CGO / no C compiler required)

	"example.com/rest-api/models"
	_ "modernc.org/sqlite"
)

// ErrEventNotFound is returned when an event lookup by id finds no row.
var ErrEventNotFound = errors.New("event not found")

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
func GetAllDbEvents() ([]models.Event, error) {
	query := `SELECT id, name, description, location, dateTime, user_id FROM events`

	rows, err := DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query events: %w", err)
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var (
			event  models.Event
			dtText string
		)
		if err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&dtText,
			&event.UserId,
		); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}

		// dateTime is stored as TEXT (RFC3339); parse it back into time.Time
		event.DateTime, err = time.Parse(time.RFC3339, dtText)
		if err != nil {
			return nil, fmt.Errorf("parse dateTime %q: %w", dtText, err)
		}

		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate events: %w", err)
	}
	return events, nil
}
func InsertData(name string, description string, location string, dateTime string, user_id int) error {
	query := `
	INSERT INTO events (
		name,
		description,
		location,
		dateTime,
		user_id
	) VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare insert: %w", err)
	}
	defer stmt.Close()

	// stmt already knows the SQL — only pass the placeholder values, in order.
	if _, err := stmt.Exec(name, description, location, dateTime, user_id); err != nil {
		return fmt.Errorf("exec insert: %w", err)
	}
	return nil
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

func GetIdDbEvents(id int) (models.Event, error) {
	query := `
		SELECT id, name, description, location, dateTime, user_id
		FROM events
		WHERE id = ?
	`

	var (
		event  models.Event
		dtText string
	)
	err := DB.QueryRow(query, id).Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&dtText,
		&event.UserId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Event{}, fmt.Errorf("%w (id=%d)", ErrEventNotFound, id)
		}
		return models.Event{}, fmt.Errorf("query event %d: %w", id, err)
	}

	// dateTime is stored as TEXT (RFC3339); parse it back into time.Time
	event.DateTime, err = time.Parse(time.RFC3339, dtText)
	if err != nil {
		return models.Event{}, fmt.Errorf("parse dateTime %q: %w", dtText, err)
	}
	return event, nil
}

func DeleteEventById(id int) error {
	query := `DELETE FROM events WHERE id = ?`
	result, err := DB.Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
