package models

import (
	"event-management-api/db"
	"time"
)

type Event struct {
	ID          int64
	NAME        string `binding: "required"`
	DESCRIPTION string `binding: "required"`
	LOCATION    string `binding: "required"`
	DATE        time.Time
	USER_ID     int
}

var events = []Event{}

func (e Event) SAVE() {
	query := `insert into events(name,description,location,date,user_id) values (?,?,?,?,?);`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		panic(err)
	}
	res, err := stmt.Exec(e.NAME, e.DESCRIPTION, e.LOCATION, e.DATE, e.USER_ID)

	if err != nil {
		panic(err)
	}
	id, err := res.LastInsertId()

	if err != nil {
		panic(err)
	}
	e.ID = id

	events = append(events, e)
}

func GetAllEvents() ([]Event, error) {
	query := "Select * from events"
	res, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer res.Close()

	var events []Event

	for res.Next() {
		var event Event
		err := res.Scan(&event.ID, &event.NAME, &event.DESCRIPTION, &event.LOCATION, &event.DATE, &event.USER_ID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	var event Event
	query := "Select * from events where id=?"
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&event.ID, &event.NAME, &event.DESCRIPTION, &event.LOCATION, &event.DATE, &event.USER_ID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}
