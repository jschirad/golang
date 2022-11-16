package main

import "fmt"

// Definimos un struct - clase
type Employee struct {
	id   int
	name string
}

// Añadimos un metodo a nuestra clase - reciber function
func (e *Employee) SetId(id int) {
	e.id = id
}

// Añadimos un metodo
func (e *Employee) SetName(name string) {
	e.name = name
}

func main() {
	// Creamos un objeto
	e := Employee{}
	fmt.Printf("%v\n", e)
	// añadimos valores
	e.id = 1
	e.name = "Juan"
	fmt.Printf("%v\n", e)
	e.SetId(5)
	e.SetName("Pedro")
	fmt.Printf("%v\n", e)
	fmt.Println(e.id)
	fmt.Println(e.name)
}
