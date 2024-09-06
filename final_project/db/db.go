package db

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	DB = db

	err = createTables()
	if err != nil {
		panic("Database could not connect: " + err.Error())
	}

	fmt.Println("Tables created successfully!")
}

func createTables() error {
	createUserTable := `
	      CREATE TABLE IF NOT EXISTS users (
	          id INTEGER PRIMARY KEY AUTOINCREMENT,
						name TEXT NOT NULL,
	          password TEXT NOT NULL,
	          email TEXT NOT NULL UNIQUE,
	          created_at DATETIME NOT NULL
	      )
	  `

	_, err := DB.Exec(createUserTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}
	createEventsTable := `
	      CREATE TABLE IF NOT EXISTS events (
	          id INTEGER PRIMARY KEY AUTOINCREMENT,
	          title TEXT NOT NULL,
	          description TEXT NOT NULL,
	          location TEXT NOT NULL,
						start_date DATETIME,
						end_date DATETIME,
	          created_at DATETIME NOT NULL,
	          user_id INTEGER,
						FOREIGN KEY(user_id) REFERENCES users(id)
	      )
	  `
	// deleteTable := `DROP TABLE IF EXISTS events`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}
	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registrations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER,
			user_id INTEGER,
			FOREIGN KEY(event_id) REFERENCES events(id),
			FOREIGN KEY(user_id) REFERENCES users(id)
		)
	`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("Could not create users table: " + err.Error())
	}
	return err
}
