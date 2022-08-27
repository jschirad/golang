package main

import "fmt"

func main() {
	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado ", areaCuadrado)

	x := 10
	y := 50

	//Suma
	result := x + y
	fmt.Println("Suma ", result)

	// Resta
	result = y - x
	fmt.Println("Resta ", result)

	// Multiplicación
	result = x * y
	fmt.Println("Multiplicación ", result)

	// División
	result = y / x
	fmt.Println("División ", result)

	// Modulo
	result = y % x
	fmt.Println("Modulo ", result)

	// Incremental
	x++
	fmt.Println("x++ ", x)

	// Decremental
	x--
	fmt.Println("x--", x)
}
