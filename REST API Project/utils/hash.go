// Package utils contains small, dependency-free helpers shared across the project.
// hash.go isolates the password-hashing algorithm (currently bcrypt) so callers
// don't have to import "golang.org/x/crypto/bcrypt" directly and so swapping the
// algorithm later (argon2, scrypt, etc.) is a single-file change.
package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns a bcrypt hash of the given plain-text password.
// The returned string is safe to store directly in the database.
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("hash password: %w", err)
	}
	return string(hashed), nil
}

// CheckPasswordHash reports whether the given plain-text password matches the
// stored bcrypt hash. It returns true only on an exact match; any error from
// bcrypt (including mismatch) is treated as "not a match".
func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
