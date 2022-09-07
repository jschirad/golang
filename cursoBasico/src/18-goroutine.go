package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(text)
}

func main() {
	// wait group
	var wg sync.WaitGroup
	// Print message
	fmt.Println("Hello")
	// agregar routine al wg
	wg.Add(1)
	// go routine
	go say("world", &wg)
	// Espera 1 go routine
	wg.Wait()

	// Funciones anonimas como routine
	go func(text string) {
		fmt.Println(text)
	}("Bye!")
	time.Sleep(time.Second * 1)
}
