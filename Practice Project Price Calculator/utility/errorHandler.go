package utility

import (
	"errors"
	"fmt"
)

func ErrorHandle(message string, printMessage ...string) error {
	if len(printMessage) > 0 {
		fmt.Println(printMessage[0])
	}
	return errors.New(message)
}
