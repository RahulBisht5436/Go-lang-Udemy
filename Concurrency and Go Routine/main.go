package main

import (
	"fmt"
	"time"
)

func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

// here inside the slow greet we get a chan param with name doneChan and type bool
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	
	// here we return the doneChan chan with value true
	doneChan <- true

}

func main() {
	// Adding this go keyword in front
	// tells go to run them as go routine(parallel execution)
	go greet("Nice to meet you!")
	go greet("How are you?")
	
	// we created a chan with name done and has type bool
	done := make(chan bool)
	
	// we have send the chan inside the function
	go slowGreet("How ... are ... you ...?", done)
	go greet("I hope you're liking the course!")

	// here we wait for the channel for responding
	<-done
}
