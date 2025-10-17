package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.Marshal(todo)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, []byte(json), 0644)
}

func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, fmt.Errorf("text cannot be empty")
	}
	return Todo{
		Text: text,
	}, nil

}
