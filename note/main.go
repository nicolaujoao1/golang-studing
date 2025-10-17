package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	displayer
}

func main() {
	title, content := getNoteData()
	todoText := getTodoData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	outputData(userNote)
	outputData(todo)
}

//	func saveData(s saver) error {
//		err := s.Save()
//		if err != nil {
//			fmt.Println("Error saving data:", err)
//			return err
//		}
//		return nil
//	}
//
//	func printSomeThing(value any) {
//		fmt.Println(value)
//	}
func outputData(o outputtable) {
	o.Display()
	err := o.Save()
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}
	fmt.Println("Data saved successfully!")
}

func getTodoData() string {
	var text, err = getUseInput("Enter todo text:")
	if err != nil {
		return ""
	}
	return text
}
func getNoteData() (string, string) {

	tittle, err := getUseInput("Note title:")

	if err != nil {
		fmt.Println("Error:", err)
		return "", ""
	}
	content, err := getUseInput("Note content:")
	if err != nil {
		fmt.Println("Error:", err)
		return "", ""
	}
	return tittle, content
}
func getUseInput(prompt string) (string, error) {

	fmt.Println(prompt)
	text, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text, nil
}
