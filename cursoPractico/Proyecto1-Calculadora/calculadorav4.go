package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct {
}

// Recivers function
func (calc) operate(entrada string, operador string) int {

	entradaLimpia := strings.Split(entrada, operador)
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])
	switch operador {
	case "+":
		fmt.Println("La suma es:", operador1+operador2)
		return operador1 + operador2
	case "-":
		fmt.Println("La resta es", operador1-operador2)
		return operador1 - operador2
	case "*":
		fmt.Println("La multiplicacón es", operador1*operador2)
		return operador1 * operador2
	default:
		fmt.Println("La operación no esta definida.")
		return 0
	}
}

// Funcion para utilizar atoi
func parsear(entrada string) int {
	operador, _ := strconv.Atoi(entrada)
	return operador
}

// Funcion para leer entrada
func leerEntrada() string {
	// scannear argumentos por pantalla
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Almacenamos el argumento como string
	return scanner.Text()
}

func main() {

	entrada := leerEntrada()
	operador := leerEntrada()
	fmt.Println(entrada)
	fmt.Println(operador)

	c := calc{}
	c.operate(entrada, operador)
}
