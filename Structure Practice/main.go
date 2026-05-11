package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	// Custom packages for Note and Todo functionality
	"example.com/strucPractice/note"
	"example.com/strucPractice/todo"
)

type saver interface {
	Save() error
}
type Displayer interface {
	Display()
	saver
}

func displayProperties(data Displayer) {
	err := saveData(data)
	if err != nil {
		return
	}
	data.Display()
}
func saveData(data saver) error {
	erro := data.Save()
	if erro != nil {
		println("Failed to save the data")
		return erro
	}
	println("data is saved SuccessFully")
	return nil
}

// main is the entry point of the program.
// It collects user input for a Note (title + content) and a Todo (text),
// then creates, displays, and saves both to their respective files.

func main() {
	// Prompt the user to enter the note's title
	title, error := getUserInput("Note Title")
	if error != nil {
		println(error.Error())
		return
	}

	// Prompt the user to enter the note's content
	content, error := getUserInput("Note Content")
	if error != nil {
		println(error.Error())
		return
	}

	// Prompt the user to enter the todo text
	text, error := getUserInput("Enter the text here")
	if error != nil {
		println(error.Error())
		return
	}

	// Create a new Todo using the text entered by the user.
	// todo.New returns a Todo struct and an error if the text is empty.
	userTodo, userTodoError := todo.New(text)
	if userTodoError != nil {
		println(userTodoError.Error())
		return
	}

	// Print the Todo details to the terminal
	displayProperties(userTodo)
	// Save the Todo to "todoData.txt"
	saveData(userTodo)
	// userTodo.Save()

	// Create a new Note using the title and content entered by the user.
	// note.New returns a Note struct and an error if either field is empty.
	userNote, userError := note.New(title, content)

	// Print the Note details to the terminal
	// userNote.Display()
	displayProperties((userNote))

	// Save the Note to "noteData.txt"
	saveData(userNote)
	// userNote.Save()

	// Check for error AFTER Display/Save — note that if userError is not nil,
	// userNote will be an empty Note{} so the display and save above would have run on empty data.
	// Ideally, this check should happen right after note.New (before Display/Save).
	if userError != nil {
		println("Error in creating the Note")
		return
	}
}

// getUserInput prints a prompt to the terminal and reads a full line of text from the user.
// It returns the input string and any error encountered during reading.
//
// Parameters:
//   - prompt: the message shown to the user (e.g., "Note Title")
//
// Returns:
//   - string: the text the user typed
//   - error:  non-nil if input was empty or a read error occurred
func getUserInput(prompt string) (string, error) {

	// Variable that will store the final user input
	var userData string

	// Print the message/prompt for the user
	// Example: "Enter Your Name"
	fmt.Println(prompt)

	// -----------------------------
	// bufio.NewReader(os.Stdin)
	// -----------------------------

	// os.Stdin
	// ----------
	// Stdin = Standard Input
	// It represents the keyboard input stream.
	// Whatever the user types in terminal comes through os.Stdin.

	// bufio.NewReader(...)
	// --------------------
	// Creates a buffered reader.
	// A buffered reader reads input more efficiently
	// and allows us to read complete lines easily.

	// Why use bufio instead of fmt.Scanln ?
	// -------------------------------------
	// fmt.Scanln stops reading at spaces.
	// Example:
	// Input: Rahul Bisht
	// fmt.Scanln => only "Rahul"
	//
	// bufio.Reader can read the entire line including spaces:
	// "Rahul Bisht"

	reader := bufio.NewReader(os.Stdin)

	// ReadString('\n')
	// ----------------
	// Reads user input until it finds a newline character '\n'
	// (i.e., until the user presses ENTER).
	//
	// Example:
	// User types: Hello World + ENTER
	// userData = "Hello World\n"
	//
	// Note: the '\n' character is included in the returned string.

	userData, err := reader.ReadString('\n')

	// Validation check
	// ----------------
	// If user entered empty data
	// OR some reading error occurred,
	// return an error.

	if userData == "" || err != nil {
		return "", errors.New("Entered the Empty or Wrong data , Only string data is allowed")
	}

	// If everything is correct,
	// return the entered user data and nil error.
	
	return userData, nil
}
func callingFunction(){
	var results = printSomethings(1, 2)
	fmt.Println(results)
}
func printSomethings[T int|float32 | float64](val1, val2 T) T {
	return val1 + val2
}
