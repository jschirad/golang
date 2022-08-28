package main

import "fmt"

func main() {
	// Defer - ultima tarea de la función
	defer fmt.Println("Adios!")
	fmt.Println("Hola Mundo!")

	// Continue
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}
		// Break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
