package main

import "fmt"

// put data into the channel
func say(text string, c chan string) {
	c <- text
}
func main() {
	// create a channel
	c := make(chan string)

	fmt.Println("Hello")
	// function
	go say("Bye", c)
	// Print the data into the channel
	fmt.Println(<-c)
}
