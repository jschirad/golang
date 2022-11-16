package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Declarando variables
	var x int
	x = 8
	y := 7
	// Imprimiendo variables
	fmt.Println(x)
	fmt.Println(y)
	// Convertimos un string a entero
	myValue, err := strconv.ParseInt("7", 0, 64)
	// myValue, err := strconv.ParseInt("NaN", 0, 64) Para encontrar un error
	// La funcion strconv nos devuelve dos variables
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Println(myValue)
	}
	// Creando un map
	m := make(map[string]int)
	m["Key"] = 6
	fmt.Println(m["Key"])
	fmt.Println(m)
	// Creando un slice
	s := []int{1, 2, 3}
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
	// Añadimos valores al slice usando append
	s = append(s, 16)
	for index, value := range s {
		fmt.Println(index)
		fmt.Println(value)
	}
}
