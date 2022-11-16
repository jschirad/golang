package main

import "fmt"

// Definimos un struct - clase
type Employee struct {
	id       int    // zero value
	name     string // zero value
	vacation bool   // default to false
}

// Simulamos un constructor con una función
func NewEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Creamos un Employee
	e := Employee{}
	// Print
	fmt.Printf("%v\n", e)

	// Creamos otro Employee - construyendo
	e2 := Employee{
		id:       1,
		name:     "Juan",
		vacation: true,
	}
	// Print
	fmt.Printf("%v\n", e2)

	// Creamos otro Employee - new devuelve un apuntador
	e3 := new(Employee)
	// Print
	fmt.Printf("%v\n", *e3)

	// Creamos otro Employee
	e4 := NewEmployee(2, "Pedro", true)
	fmt.Printf("%v\n", *e4)
}
