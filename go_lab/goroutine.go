package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a channel for synchronization
	message := make(chan string)

	// Launch a goroutine to send a message to the channel after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		message <- "Hello, World from a goroutine!"
	}()

	// Launch another goroutine with a timer that changes the message
	go func() {
		for i := 3; i >= 1; i-- {
			fmt.Printf("Starting in %d seconds!\n", i)
			time.Sleep(1 * time.Second)
		}
		message <- "BOOM! Here's your final 'Hello, World!'"
	}()

	// Receive the message from the channel and print it
	fmt.Println(<-message)
	fmt.Println(<-message)
}
