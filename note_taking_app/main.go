package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/go-project/note"
	"example.com/go-project/todo"
)

// interface uses type keyword
// common convinsion interface has -er suffix
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// embedded interface
// common convinsion interface has -er suffix
type outputable interface {
	saver // you can inherit from multiple interfaces
	Display()
}

func main() {
	title, content := getUserNote()
	todoText := getUserInput("Enter Todo: ")
	todo, err := todo.New(todoText)

	if err != nil {
		fmt.Println(err)
		return
	}

	err = outputData(todo)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(userNote)

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

func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		fmt.Println("Saving Failed", err)
		return err
	}
	fmt.Println("Saved!")
	return nil
}

// this interface takes both note and todo
func outputData(data outputable) error {
	data.Display()

	err := saveData(data)
	if err != nil {
		fmt.Println("Saving Failed", err)
		return err
	}

	fmt.Println("Saved!")
	return nil
}

// wild card: value can be anything just like any type in ts
func printSomething(value interface{}) interface{} {
	// typedValue: int value
	// ok: bool if it is an int
	intVal, ok := value.(int)
	if ok {
		intVal += 1
		return intVal
	}
	float64Val, ok := value.(float64)
	if ok {
		float64Val += 1
		return float64Val
	}
	strVal, ok := value.(string)
	if ok {
		strVal = strings.ToUpper(strVal)
		return strVal
	}
	// type checking with switch
	// switch value.(type) {
	// case string:
	// 	fmt.Println("string:", value)
	// case int:
	// 	fmt.Println("int:", value)
	// case float64:
	// 	fmt.Println("float64:", value)
	// default:
	// 	fmt.Println("unknown type")
	// }
	// fmt.Println(value)
	return nil
}

// Generic function
// T: generic type
// you can use any as interface{}
// you could add more generic types with ,
// func add[T any, K any](a K,b T ) (T) {
// you could specify options of T
func add[T int | float64 | string](a, b T) T {
	return a + b
}
