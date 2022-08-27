package main

import "fmt"

func main() {
	// Ciclo for conditional
	fmt.Println("For Conditional")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	// Ciclo while
	fmt.Println("While Conditional")
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}
	fmt.Println("For Forever Conditional")
	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}
