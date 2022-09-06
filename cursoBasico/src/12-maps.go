package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepito"] = 20

	fmt.Println(m)

	// recorremos un map con range
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontramos valor
	value := m["Jose"]
	fmt.Println(value)

	// Imprimir valor inexistente
	value1 := m["Juan"]
	fmt.Println(value1)

	value2, ok := m["Julio"] // ok Comprueba si existe
	fmt.Println(value2, ok)

	value3, ok := m["Pepito"]
	fmt.Println(value3, ok)
}
