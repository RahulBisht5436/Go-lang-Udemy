package models

import "time"

// Event represents one row in the `events` table.
//
// UserId / UserEmail are NOT taken from the request body — they are filled
// in from the authenticated JWT in the route handler. That's why they have
// no `binding:"required"` tag: a client trying to pass them in is simply
// ignored / overwritten.
type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"        binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int64     `json:"userId"`
	UserEmail   string    `json:"userEmail"`
}
