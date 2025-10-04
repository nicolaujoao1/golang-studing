package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Note created successfully!")
	userNote.Display()
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
