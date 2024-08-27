package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
)

func main() {
	title, content := getUserNote()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving Note Failed", err)
		return
	}
	fmt.Println("Note Saved")
}
func getUserNote() (string, string) {
	title := getUserInput("Enter title: ")

	content := getUserInput("Enter content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	// var value string
	// you cannot get long input with Scanln
	// fmt.Scanln(&value)

	// you can get long input with bufio with os.Stdin
	reader := bufio.NewReader(os.Stdin)
	// stop reading at new line use
	// single quote for single char double quote for string
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	// remove specific char from end
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
