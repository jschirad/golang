package main

// En esta clase creamos mypackage
import (
	pk "cursoBasico/src/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2022
	fmt.Println(myCar)

	// var mySecondCar pk.carPrivate
	// fmt.Println(mySecondCar)

	// Public function
	pk.PrintMessage("Hola Mundo!")
	// Private function
	pk.printMessage1("Hola!")
}
