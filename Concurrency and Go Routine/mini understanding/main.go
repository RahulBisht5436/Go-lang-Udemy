package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

// here inside the slow greet we get a chan param with name doneChan and type bool
func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)

	// here we return the doneChan chan with value true
	doneChan <- true

}

func main() {

	// making a array of chan in order to wait for all the chan
	dones := make([]chan bool, 4)

	// we created a chan with name done and has type bool
	for index, _ := range dones {
		dones[index] = make(chan bool)
	}
	// Adding this go keyword in front
	// tells go to run them as go routine(parallel execution)
	go greet("Nice to meet you!", dones[0])

	// we are passing chan inside the dones chan bool array
	go greet("How are you?", dones[1])

	// we have send the chan inside the function
	go slowGreet("How ... are ... you ...?", dones[2])
	go greet("I hope you're liking the course!", dones[3])

	// here we wait for the channel for responding
	//using the for loop we wait for each Chan
	for _, value := range dones {
		<-value
	}

}
