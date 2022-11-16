package main

import "fmt"

// Definimos un struct - clase
type Employee struct {
	id   int
	name string
}

func main() {
	// Creamos un objeto
	e := Employee{}
	fmt.Printf("%v\n", e)
	// añadimos valores
	e.id = 1
	e.name = "Juan"
	fmt.Printf("%v\n", e)
}
