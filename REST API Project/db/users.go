package db

import (
	"database/sql"
	"errors"
	"fmt"

	"example.com/rest-api/models"
	"example.com/rest-api/utils"
	"modernc.org/sqlite"
)

// ErrUserEmailExists is returned when a signup attempt uses an already-registered email.
var ErrUserEmailExists = errors.New("email already registered")

// ErrInvalidCredentials is returned when an email isn't found OR the password doesn't match.
// We deliberately return the same error in both cases to avoid leaking which emails
// are registered (account enumeration prevention).
var ErrInvalidCredentials = errors.New("invalid credentials")

func SaveUser(u *models.User) error {
	hashed, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare(`INSERT INTO users(email, password) VALUES (?, ?)`)
	if err != nil {
		return fmt.Errorf("prepare insert user: %w", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(u.Email, hashed)
	if err != nil {
		// Translate the UNIQUE constraint violation into a domain error
		// so callers can return HTTP 409 instead of a generic 500.
		var sqliteErr *sqlite.Error
		if errors.As(err, &sqliteErr) && sqliteErr.Code() == 2067 {
			return ErrUserEmailExists
		}
		return fmt.Errorf("exec insert user: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("last insert id: %w", err)
	}

	u.ID = id
	u.Password = "" // never let the hashed/plain password leave this layer
	return nil
}

// AuthenticateUser looks up the user by email and verifies the supplied password
// against the stored bcrypt hash. On success it returns the user with the password
// field cleared. On any failure (no such email OR wrong password) it returns
// ErrInvalidCredentials.
func AuthenticateUser(email, password string) (models.User, error) {
	var (
		user       models.User
		hashedPass string
	)

	err := DB.QueryRow(
		`SELECT Id, email, password FROM users WHERE email = ?`,
		email,
	).Scan(&user.ID, &user.Email, &hashedPass)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.User{}, ErrInvalidCredentials
		}
		return models.User{}, fmt.Errorf("query user: %w", err)
	}

	if !utils.CheckPasswordHash(password, hashedPass) {
		return models.User{}, ErrInvalidCredentials
	}

	return user, nil
}
