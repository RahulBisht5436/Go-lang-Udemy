package note

import (
	"errors"
	"fmt"
	"os"
	"time"
)

type Note struct {
	title     string `json:"title_on"`
	content   string `json:updated_content`
	createdAt time.Time
}

func (n Note) Display() {
	fmt.Println("This is the title : ")
	fmt.Println(n.title)
	fmt.Println("This is the content : ")
	fmt.Println(n.content)
}
func (n Note) Save() error {
	text := fmt.Sprintf("Title: %v\nContent: %v\nCreated At: %v", n.title, n.content, n.createdAt)
	return os.WriteFile("noteData.txt", []byte(text), 0644)
}
func New(title string, content string) (Note, error) {
	if title == "" || content == "" {
		return Note{}, errors.New("Invalid or insuffcient Data")
	}
	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
