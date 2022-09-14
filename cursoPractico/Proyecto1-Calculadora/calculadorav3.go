package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// scannear argumentos por pantalla
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// Almacenamos el argumento como string
	operacion := scanner.Text()
	fmt.Println(operacion)
	operador := "-"
	valores := strings.Split(operacion, operador)
	fmt.Println(valores)
	// Al imprimir los valores podemos ver que siguen siendo caracteres y no enteros.
	fmt.Println(valores[0] + valores[1])

	// Convertimos los caracteres a enteros con Atoi
	operador1, err1 := strconv.Atoi(valores[0])
	if err1 != nil {
		fmt.Println("ERROR: ", err1)
	} else {
		fmt.Println(operador1)
	}
	operador2, _ := strconv.Atoi(valores[1])
	// var suma int
	// suma = operador1 + operador2
	// fmt.Println("La suma es", suma)

	// Gestion de errores
	// Si introduciomos valores que no pueden ser convertidos a enteros, tendremos un error
	// fmt.Println(operador1)
	// El cero que obtenemos representa que hay un error
	// fmt.Println(err1) //strconv.Atoi: parsing "2adas": invalid syntax

	// Para poder manejar varias operaciones, vamos a implementar switches
	switch operador {
	case "+":
		fmt.Println("La suma es:", operador1+operador2)
	case "-":
		fmt.Println("La resta es", operador1-operador2)
	case "*":
		fmt.Println("La multiplicacón es", operador1*operador2)
	default:
		fmt.Println("La operación no esta definida.")
	}
}
