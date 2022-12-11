package main

import "fmt"

func main() {
	// Declarate const variables
	const pi float64 = 3.14
	const pi2 = 3.1415

	fmt.Println("pi ", pi)
	fmt.Println("pi2", pi2)

	// Daclarate int variables
	base := 12
	var altura int = 14
	var area int

	fmt.Println(area, altura, base)

	// Declarate zero variables
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Square area
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado ", areaCuadrado)
}
