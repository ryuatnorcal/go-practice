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

func GetAllEvents() []Event {
	return events
}
