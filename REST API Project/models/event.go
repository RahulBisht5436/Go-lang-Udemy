package models

import "time"

type Event struct {
	ID          int    `binding:"required`
	Name        string `binding:"required`
	Description string
	Location    string
	DateTime    time.Time
	UserId      int `binding:"required`
}

var events = []Event{}

func (E Event) Save() {
	// add to the Database later
	events = append(events, E)
}

func GetAllEvent() []Event {
	return events
}
