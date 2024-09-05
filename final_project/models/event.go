package models

import (
	"time"

	"example.com/final_project/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"` // binding required for require from request
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	CreatedAt   time.Time `json:"created_at"`
	UserId      int       `json:"user_id"`
}

var events = []Event{}

func (e Event) Save() error {
	query := `INSERT INTO events (title, description, location, start_date, end_date, created_at, user_id)
	VALUES (?, ?, ?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Exec: Query for modify data
	result, err := stmt.Exec(e.Title, e.Description, e.Location, e.StartDate, e.EndDate, e.CreatedAt, e.UserId)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	// later add DB to store
	// events = append(events, e)
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	// Query for many rows
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	// next return true as long as there is next data in rows
	for rows.Next() {
		var event Event
		if err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.StartDate, &event.EndDate, &event.CreatedAt, &event.UserId); err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	// QueryRow: Query for one row
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.StartDate, &event.EndDate, &event.CreatedAt, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := "UPDATE events SET title = ?, description = ?, location = ?, start_date = ?, end_date = ?, created_at = ?, user_id = ? WHERE id = ?"

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Title, event.Description, event.Location, event.StartDate, event.EndDate, event.CreatedAt, event.UserId, event.ID)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.ID)
	return err
}
