package mypackage

import "fmt"

// Con mayus acceso publico
type CarPublic struct {
	Brand string
	Year  int
}

// con minus acceso privado
type carPrivate struct {
	brand string
	year  int
}

// print mensaje
func PrintMessage(text string) {
	fmt.Println(text)
}

func printMessage1(text string) {
	fmt.Println(text)
}
