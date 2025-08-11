package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.DB")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUsersTable := ` CREATE TABLE IF NOT EXISTS USERS (
		ID INTEGER PRIMARY KEY AUTOINCREMENT,
		EMAIL TEXT NOT NULL,
		PASSWORD TEXT NOT NULL
	);`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("could not create user table: %v", err)
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS EVENTS 
							(ID INTEGER PRIMARY KEY AUTOINCREMENT, NAME TEXT NOT NULL, DESCRIPTION TEXT NOT NULL, LOCATION TEXT NOT NULL,  DATE DATETIME NOT NULL, USER_ID INTEGER, FOREIGN KEY(USER_ID) REFERENCES USERs(ID) ON DELETE CASCADE);`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("could not create event table")
	}

	createRegistrationTable := `CREATE TABLE IF NOT EXISTS REGISTRATIONS
								(ID INTEGER PRIMARY KEY AUTOINCREMENT,EVENT_ID INTEGER, USER_ID INTEGER, FOREIGN KEY(EVENT_ID) REFERENCES EVENTS(ID), FOREIGN KEY(USER_ID) REFERENCES USERS(ID));`
	_, err = DB.Exec(createRegistrationTable)
}
