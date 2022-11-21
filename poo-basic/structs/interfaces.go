package main

import "fmt"

// Definimos un struct - clase
type Person struct {
	name string
	age  int
}

// Definimos la clase Empleado
type Employee struct {
	id int
}

// Definimos la clase FullTime que hereda los atributos de Persona y Empleado
type FullTimeEmployee struct {
	Person
	Employee
}

func (ftEmployee FullTimeEmployee) getMessage() string {
	return "Full Time Employee"
}

type TemporalTimeEmployee struct {
	Person
	Employee
	taxRate int
}

func (tEmployee TemporalTimeEmployee) getMessage() string {
	return "Temporal Time Employee"
}

type PrintInfo interface {
	getMessage() string
}

func getMessage(p PrintInfo) {
	fmt.Println(p.getMessage())
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Juan"
	ftEmployee.age = 26
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)

	tEmployee := TemporalTimeEmployee{}
	getMessage(tEmployee)
	getMessage(ftEmployee)

}
