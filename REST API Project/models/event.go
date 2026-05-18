package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"        binding:"required"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int       `json:"userId"      binding:"required"`
}
