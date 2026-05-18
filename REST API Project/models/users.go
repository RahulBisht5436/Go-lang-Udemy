package models

// User is the canonical user record. The strict validation rules on Email and
// Password are intended for signup; login uses LoginRequest below so it can
// accept any well-formed JSON without rejecting malformed emails outright
// (a wrong email should produce 401 Invalid credentials, not 400).
type User struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

// LoginRequest is the bind target for the /login route. We only require that
// the two fields are present; the credentials themselves are validated by
// looking them up in the database.
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
