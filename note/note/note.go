package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	tittle    string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Println("Display Note")
	fmt.Println("Title:", note.tittle)
	fmt.Println("Content:", note.content)
	fmt.Println("Created At:", note.createdAt.Format(time.RFC1123))
}
func New(tittle, content string) (Note, error) {
	if tittle == "" || content == "" {
		return Note{}, errors.New("tittle and content cannot be empty")
	}
	return Note{
		tittle:    tittle,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
