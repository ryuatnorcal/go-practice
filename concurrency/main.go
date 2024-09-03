package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChannel chan bool) {
	fmt.Println("Hello!", phrase)
	doneChannel <- true
}

func slowGreet(phrase string, doneChannel chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChannel <- true // signal that the task is done
	// once close channe, you cannot send data anymore
	close(doneChannel)
}

// adding go keyword runs parallel
// channel
func main() {
	//channel: gives status of go routine
	dones := make([]chan bool, 4)
	dones[0] = make(chan bool)
	go greet("Nice to meet you!", dones[0])
	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])
	dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[2])
	dones[3] = make(chan bool)
	go greet("I hope you're liking the course!", dones[3])
	// wait for all go routines to finish
	// <-done
	// <-done
	// <-done
	// <-done
	for _, done := range dones {
		<-done
	}

	// for _, done := range dones {
	// 	fmt.Println(done)
	// }
}
