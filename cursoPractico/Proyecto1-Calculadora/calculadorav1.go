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
	valores := strings.Split(operacion, "+")
	fmt.Println(valores)
	// Al imprimir los valores podemos ver que siguen siendo caracteres y no enteros.
	fmt.Println(valores[0] + valores[1])

	// Convertimos los caracteres a enteros con Atoi
	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])
	var suma int
	suma = operador1 + operador2
	fmt.Println("La suma es", suma)
}
