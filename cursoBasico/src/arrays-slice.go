package main

import "fmt"

func main() {
	// Array
	fmt.Println("Arrays")
	var array [4]int
	array[0] = 1
	array[1] = 2
	// Array - Lenght - capacity
	fmt.Println(array, len(array), cap(array))

	// Slice
	fmt.Println("Slices")
	slice := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(slice, len(slice), cap(slice))

	// Metodos
	fmt.Println("Metodos")
	fmt.Println(slice[0])
	fmt.Println(slice[:3])
	fmt.Println(slice[2:4])
	fmt.Println(slice[4:])

	// Append
	fmt.Println("Append")
	slice = append(slice, 6)
	fmt.Println(slice)

	// Append New list
	fmt.Println("List + List")
	newSlice := []int{7, 8, 9, 10}
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}
