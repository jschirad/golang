package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

// Modificar output de mi struct
func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disk y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, disk: 512, brand: "apple"}

	fmt.Println(myPc)
}
