package main

import "fmt"

func main() {
	// Declaración de mensajes
	helloMessage := "Hello"
	worldMessage := "world!"

	// Print ln - Agrega salto de linea
	fmt.Println(helloMessage, worldMessage)

	// Print f - Agrega función
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos", nombre, cursos)
	fmt.Println(message)

	// Tipos de Datos
	fmt.Printf("helloMessage %T\n", helloMessage)
	fmt.Printf("Cursos %T\n", cursos)
}
