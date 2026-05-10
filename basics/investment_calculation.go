package main

//this is how we can have multiple  inbuild imports
import (
	"fmt"
)

func main() {
	var messageObject retrunType
	var newMessageFormatted string
	messageObject, newMessageFormatted = WrapperFunction("Hello nice to meet you")
	fmt.Printf(newMessageFormatted, messageObject)
}

type retrunType struct {
	message    string
	code       string
	statusCode int64
}

func WrapperFunction(message string) (retrunType, string) {
	fmt.Println(message)
	var newmessageString string = message + "this is the returned value from the function"
	return retrunType{
		message:    newmessageString,
		code:       "successful",
		statusCode: 200,
	}, "this is the second returned"

}
