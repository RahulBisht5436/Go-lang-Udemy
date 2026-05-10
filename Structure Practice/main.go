package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"example.com/strucPractice/note"
)


func main() {
	title, error := getUserInput("Note Title")
	if error != nil {
		println(error.Error())
		return
	}
	content, error := getUserInput("Note Content")
	if error != nil {
		println(error.Error())
		return
	}
	userNote, userError := note.New(title, content)
	userNote.DisplayNote()
	userNote.Save()
	if userError != nil {
		println("Error in creating the Note")
		return
	}
}

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
	// bufio.Reader can read the entire line:
	// "Rahul Bisht"

	reader := bufio.NewReader(os.Stdin)

	// ReadString('\n')
	// ----------------
	// Reads user input until it finds a newline character '\n'
	//
	// Example:
	// User types:
	// Hello World + ENTER
	//
	// It reads everything until ENTER is pressed.

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
	// return the entered user data
	// and nil error.
	return userData, nil
}
