package main

import (
	"fmt"
	"strings"
)

func isPalindromo(text string) {
	var textReverse string
	minusText := strings.ToLower(text)

	for i := len(minusText) - 1; i >= 0; i-- {
		textReverse += string(minusText[i])
	}
	if minusText == textReverse {
		fmt.Println("Es Palindrome")
	} else {
		fmt.Println("No es Palindromo")
	}

}
func main() {
	slice := []string{"Hola", "que", "hace"}

	// Rango del slice
	for i := range slice {
		fmt.Println(i)
	}
	// Recorrer slice
	isPalindromo("amor a roma")
}
