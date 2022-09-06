package main

import "fmt"

// Declarando estructuras - clases
type car struct {
	brand string
	year  int
}

func main() {
	myCar := car{brand: "Ford", year: 2022}
	fmt.Println(myCar)

	// otra manera de crear clase
	var mySecondCar car
	mySecondCar.brand = "Ferrari"
	fmt.Println(mySecondCar)
}
