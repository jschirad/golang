package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) ping() {
	fmt.Println(myPc.brand, "Pong")
}
func (myPc *pc) duplicateRAM() { // Vamos a acceder a los valores mediante el puntero
	myPc.ram = myPc.ram * 2
	fmt.Println(myPc.ram)
}

func main() {

	a := 50 // declaramos un entero
	b := &a // direccion de memoria del entero

	fmt.Println(a)
	fmt.Println(b)
	// Para acceder a la direccion de memoria
	fmt.Println(*b)
	// Si modificamos la direccion de memoria, modificamos la variable inicial
	*b = 100
	fmt.Println(a)

	myPc := pc{ram: 32, disk: 512, brand: "Apple"}
	fmt.Println(myPc)

	myPc.ping()

	myPc.duplicateRAM()

	fmt.Println(myPc)
}
