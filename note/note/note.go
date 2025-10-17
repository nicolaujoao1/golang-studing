package note

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Tittle    string    `json:"tittle"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Tittle, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	os.WriteFile(fileName, []byte(json), 0644)

	return nil
}
func (note Note) Display() {
	fmt.Println("Display Note")
	fmt.Println("Title:", note.Tittle)
	fmt.Println("Content:", note.Content)
	fmt.Println("Created At:", note.CreatedAt.Format(time.RFC1123))
}
func New(tittle, content string) (Note, error) {
	if tittle == "" || content == "" {
		return Note{}, errors.New("tittle and content cannot be empty")
	}
	return Note{
		Tittle:    tittle,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}
