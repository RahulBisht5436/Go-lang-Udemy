package db

import (
	// Go uses this SQL package
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	// Pure-Go SQLite driver (no CGO / no C compiler required)

	"example.com/rest-api/models"
	_ "modernc.org/sqlite"
)

// ErrEventNotFound is returned when an event lookup by id finds no row.
var ErrEventNotFound = errors.New("event not found")

// ErrForbidden is returned when the caller is authenticated but is not the
// owner of the resource they're trying to modify. The route layer translates
// this into HTTP 403.
var ErrForbidden = errors.New("forbidden")

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
	query := `SELECT id, name, description, location, dateTime, user_id, user_email FROM events`

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
			// user_email is nullable for legacy rows created before the
			// column existed, so scan into sql.NullString to avoid a
			// "converting NULL to string is unsupported" error.
			email sql.NullString
		)
		if err := rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&dtText,
			&event.UserId,
			&email,
		); err != nil {
			return nil, fmt.Errorf("scan event: %w", err)
		}
		event.UserEmail = email.String

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

// InsertData persists a new event. The caller is responsible for setting
// event.UserId and event.UserEmail from the authenticated JWT — they MUST
// NOT come from the request body, otherwise any client could claim to own
// any event. On success the new row id is written back into event.ID.
func InsertData(event *models.Event) error {
	query := `
	INSERT INTO events (
		name,
		description,
		location,
		dateTime,
		user_id,
		user_email
	) VALUES (?, ?, ?, ?, ?, ?)
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare insert: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime.Format(time.RFC3339),
		event.UserId,
		event.UserEmail,
	)
	if err != nil {
		return fmt.Errorf("exec insert: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("last insert id: %w", err)
	}
	event.ID = int(id)
	return nil
}
func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL
	) 
	`
	_, errUserTable := DB.Exec(createUsersTable)
	if errUserTable != nil {
		panic("Not able to create the User Table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	location TEXT NOT NULL,
	dateTime TEXT NOT NULL,
	user_id INTEGER,
	user_email TEXT,
	FOREIGN KEY(user_id) REFERENCES users(Id)
	) 
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		panic("Not able to create the table: " + err.Error())
	}

	// Lightweight migration: if the events table existed BEFORE we added
	// the user_email column, CREATE TABLE IF NOT EXISTS won't touch it.
	// Try to add the column and swallow the "duplicate column" error that
	// SQLite raises on subsequent runs.
	if _, err := DB.Exec(`ALTER TABLE events ADD COLUMN user_email TEXT`); err != nil &&
		!strings.Contains(err.Error(), "duplicate column name") {
		panic("Failed to add user_email column: " + err.Error())
	}
}

func GetIdDbEvents(id int) (models.Event, error) {
	query := `
		SELECT id, name, description, location, dateTime, user_id, user_email
		FROM events
		WHERE id = ?
	`

	var (
		event  models.Event
		dtText string
		email  sql.NullString
	)
	err := DB.QueryRow(query, id).Scan(
		&event.ID,
		&event.Name,
		&event.Description,
		&event.Location,
		&dtText,
		&event.UserId,
		&email,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Event{}, fmt.Errorf("%w (id=%d)", ErrEventNotFound, id)
		}
		return models.Event{}, fmt.Errorf("query event %d: %w", id, err)
	}
	event.UserEmail = email.String

	// dateTime is stored as TEXT (RFC3339); parse it back into time.Time
	event.DateTime, err = time.Parse(time.RFC3339, dtText)
	if err != nil {
		return models.Event{}, fmt.Errorf("parse dateTime %q: %w", dtText, err)
	}
	return event, nil
}

// DeleteEventById removes an event only if it belongs to callerUserId.
// Returns ErrEventNotFound if the row doesn't exist, or ErrForbidden if
// the row exists but is owned by a different user.
func DeleteEventById(id int, callerUserId int64) error {
	ownerId, err := eventOwner(id)
	if err != nil {
		return err
	}
	if ownerId != callerUserId {
		return ErrForbidden
	}

	if _, err := DB.Exec(`DELETE FROM events WHERE id = ?`, id); err != nil {
		return fmt.Errorf("exec delete: %w", err)
	}
	return nil
}

// UpdateEvent applies the supplied changes to event `id`, but only if that
// event is owned by callerUserId. Returns ErrEventNotFound / ErrForbidden
// in the same way DeleteEventById does. user_id / user_email on the row
// are preserved — only the user-editable fields are touched.
func UpdateEvent(id int, event models.Event, callerUserId int64) error {
	ownerId, err := eventOwner(id)
	if err != nil {
		return err
	}
	if ownerId != callerUserId {
		return ErrForbidden
	}

	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`
	stmt, err := DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare update: %w", err)
	}
	defer stmt.Close()

	if _, err := stmt.Exec(
		event.Name,
		event.Description,
		event.Location,
		event.DateTime.Format(time.RFC3339),
		id,
	); err != nil {
		return fmt.Errorf("exec update: %w", err)
	}
	return nil
}

// eventOwner returns the user_id that owns event `id`, or ErrEventNotFound
// if no such event exists. Kept private — handlers should go through
// UpdateEvent / DeleteEventById which call it internally.
func eventOwner(id int) (int64, error) {
	var ownerId int64
	err := DB.QueryRow(`SELECT user_id FROM events WHERE id = ?`, id).Scan(&ownerId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, fmt.Errorf("%w (id=%d)", ErrEventNotFound, id)
		}
		return 0, fmt.Errorf("query event owner %d: %w", id, err)
	}
	return ownerId, nil
}
