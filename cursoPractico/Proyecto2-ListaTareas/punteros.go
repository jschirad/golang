package main

import "fmt"

func main() {
	x := 25
	fmt.Println(&x)
	y := &x
	fmt.Println(y)
	fmt.Println(*y)
	changeValue(x)
	fmt.Println(x)
}

func changeValue(y int) {
	y = 35
}
