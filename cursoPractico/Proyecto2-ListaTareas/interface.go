package main

import "fmt"

// Interface
type animal interface {
	mover() string
}

// No Interface
type perro struct {
}

type pez struct {
}

type pajaro struct {
}

//	func (perro) caminar() string {
//		return "Soy un perro y estoy camiando."
//	}
func (perro) mover() string {
	return "Soy un perro y estoy camiando."
}

//	func (pez) nadar() string {
//		return "Soy un pez y estoy nadando"
//	}
func (pez) mover() string {
	return "Soy un pez y estoy nadando."
}

//	func (pajaro) volar() string {
//		return "Soy un pajara y estoy volando"
//	}
func (pajaro) mover() string {
	return "Soy un pajara y estoy volando."
}

// func moverPerro(p perro) {
// 	fmt.Println(p.caminar())
// }

// func moverPez(p pez) {
// 	fmt.Println(p.nadar())
// }

// func moverPajaro(p pajaro) {
// 	fmt.Println(p.volar())
// }

func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	p := perro{}
	moverAnimal(p)
	pe := pez{}
	moverAnimal(pe)
	pa := pajaro{}
	moverAnimal(pa)
}
