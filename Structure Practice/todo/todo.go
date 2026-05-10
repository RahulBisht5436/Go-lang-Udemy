package todo

import (
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	text string `json:"title_on"`
}

func (t Todo) DisplayNote() {
	fmt.Println("This is the title : ")
	fmt.Println(t.text)
}
func (t Todo) Save() error {
	text := fmt.Sprintf("Text: %v\n", t.text)
	return os.WriteFile("todoData.txt", []byte(text), 0644)
}
func New(text string) (Todo, error) {
	if text == "" {
		return Todo{}, errors.New("Invalid or insuffcient Data")
	}
	return Todo{
		text: text,
	}, nil
}
