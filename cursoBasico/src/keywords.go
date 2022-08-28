package main

import "fmt"

func main() {
	// Defer - ultima tarea de la función
	defer fmt.Println("Hola!")
	fmt.Println("Mundo!")
}
