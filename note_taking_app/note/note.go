package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// public value only export to json
// struct tag is for json key fields only
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Display() {
	fmt.Println("Title: ", note.Title)
	fmt.Println("Content: ", note.Content)
	fmt.Println("Created at: ", note.CreatedAt)
}
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_") + ".json"
	fileName = strings.ToLower(fileName)
	json, err := json.Marshal(note)
	if err != nil {
		return err
	}
	return os.WriteFile(fileName, json, 0644)
}
func New(title string, content string) (*Note, error) {
	if title == "" || content == "" {
		return &Note{}, errors.New("Title and content are required")
	}
	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
