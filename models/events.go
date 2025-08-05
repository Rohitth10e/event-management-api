package models

import "time"

type Event struct {
	ID          int
	NAME        string `binding: "required"`
	DESCRIPTION string `binding: "required"`
	LOCATION    string `binding: "required"`
	DATE        time.Time
	USER_ID     int
}

var events = []Event{}

func (e Event) SAVE(){
	events = append(events, e)
}

func GetAllEvents() []Event{
	return events
}