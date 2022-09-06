package main

import "fmt"

func main() {
	// Para cuando quiero iterar sobre una misma variable
	modulo := 4 % 2
	switch modulo {
	case 0:
		fmt.Println("Es par")
	default:
		fmt.Println("Es impar")
	}

	// Sin condicion puede servir para anidar multiples condiciones
	value := 200
	switch {
	case value > 100:
		fmt.Println("Mayor a 100")
	case value < 0:
		fmt.Println("Menor a 0")
	default:
		fmt.Println("No condition")
	}
}
