package main

import "fmt"

type task struct {
	nombre      string
	description string
	completed   bool
}

func (t *task) marcarCompleta() {
	t.completed = true
}

func (t *task) newDescription(description string) {
	t.description = description
}

func (t *task) newName(nombre string) {
	t.nombre = nombre
}

func main() {
	t := &task{
		nombre:      "Curso Practico de Go",
		description: "Aprende Golang con mini proyectos",
		completed:   false,
	}

	fmt.Printf("%+v\n", t)

	t.marcarCompleta()
	t.newName("Curso Practico de Go Completo")
	t.newDescription("Proyectos terminados")
	fmt.Printf("%+v\n", t)

}
