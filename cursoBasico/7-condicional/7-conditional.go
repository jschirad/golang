package main

import (
	"fmt"
	"log"
	"strconv"
)

// https://pkg.go.dev/strconv
func main() {
	valor1 := 1
	valor2 := 2

	if valor1 == 1 {
		fmt.Println("El valor es 1.")
	} else {
		fmt.Println("El valor es distinto de 1.")
	}
	// if and
	if valor1 == 1 && valor2 == 2 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
	// if or
	if valor1 == 2 || valor2 == 1 {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// Conversión Texto a Numero
	value, err := strconv.Atoi("53")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Value ", value)

	value2, err2 := strconv.Atoi("hola")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("Value ", value2)
}
