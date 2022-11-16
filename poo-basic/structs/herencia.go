package main

import "fmt"

// Definimos un struct - clase
type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d \n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Juan"
	ftEmployee.age = 26
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	GetMessage(ftEmployee.Person)
}
