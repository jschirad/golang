package main

import "fmt"

func main() {
	// maps -> map[key-type]value-type
	mapa := make(map[string]int)
	mapa["a"] = 8
	fmt.Println(mapa["a"])
	fmt.Println(mapa)
}
