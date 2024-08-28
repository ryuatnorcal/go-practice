package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// public value only export to json
// struct tag is for json key fields only
type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println("Text: ", todo.Text)
}
func (todo Todo) Save() error {
	todoFile := "todo.json"

	json, err := json.Marshal(todo)
	if err != nil {
		return err
	}
	return os.WriteFile(todoFile, json, 0644)
}
func New(todo string) (*Todo, error) {
	if todo == "" {
		return &Todo{}, errors.New("Title and content are required")
	}
	return &Todo{
		Text: todo,
	}, nil
}
