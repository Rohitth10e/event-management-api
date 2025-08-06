package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS EVENTS 
							(ID INTEGER PRIMARY KEY AUTOINCREMENT, NAME TEXT NOT NULL, DESCRIPTION TEXT NOT NULL, LOCATION TEXT NOT NULL,  DATE DATETIME NOT NULL, USER_ID INTEGER);`

	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create event table")
	}

}
