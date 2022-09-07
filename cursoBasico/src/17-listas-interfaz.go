package main

import "fmt"

// Interfaz
type figuras2D interface {
	area() float64
}

// Clase cuadrado
type cuadrado struct {
	base float64
}

// Clase rectangulo
type rectangulo struct {
	base   float64
	altura float64
}

// Calculamos area cuadrado
func (c cuadrado) area() float64 {
	return c.base * c.base
}

// Calculamos area rectangulo
func (r rectangulo) area() float64 {
	return r.base * r.altura
}

// Stringer para modificar output
func (myCuadrado cuadrado) String() string {
	return fmt.Sprintf("Mi cuadrado es de %0.1f cm de lado", myCuadrado.base)
}

func (myRectangulo rectangulo) String() string {
	return fmt.Sprintf("Mi rectangulo tiene %0.1f cm de Base y %0.1f cm de Altura", myRectangulo.base, myRectangulo.altura)
}

// Funcion utilizando interfaz
func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

// Main
func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	fmt.Println(myCuadrado)
	fmt.Println(myRectangulo)

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista de interfaces
	// Emulamos un array de multiples tipos de variables
	myInterface := []interface{}{"Hola", 1, true}
	fmt.Println(myInterface...)
}
