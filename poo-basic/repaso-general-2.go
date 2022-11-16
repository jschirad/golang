package main

import (
	"fmt"
	"time"
)

func doSomething(c chan int) {
	time.Sleep(3 * time.Second)
	fmt.Println("Done")
	c <- 1
}

func main() {

	// Creamos un canal
	c := make(chan int)

	// Go Routine
	go doSomething(c)

	<-c

	// Punteros

	g := 25
	fmt.Println(g)
	h := &g
	fmt.Println(h)
	fmt.Println(*h)
}
