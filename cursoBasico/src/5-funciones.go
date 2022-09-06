package main

import "fmt"

// Imprimer mensaje
func normalFunction(message string) {
	fmt.Println(message)
}

// Imprimir argumentos
// func tripeArgument(a,b int, c string)
func tripeArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

// Retornar un valor
func returnValue(a int) int {
	return a * 2
}

// Retornar multiples valores
func doubleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	// Imprimir Mensaje
	normalFunction("Hola Mundo!")
	// Imprimir argumentos
	tripeArgument(1, 2, "Hola!")
	// Retorno de valor
	value := returnValue(2)
	fmt.Println("Value ", value)
	// Retorno de dos valores
	value1, value2 := doubleReturn(2)
	fmt.Printf("Value1 = %v Value2 = %v ", value1, value2)
	// Retorno simple
	_, value3 := doubleReturn(3)
	fmt.Printf("Value3 = %v ", value3)
	// Retorno simple
	value4, _ := doubleReturn(3)
	fmt.Printf("Value4 = %v ", value4)
}
