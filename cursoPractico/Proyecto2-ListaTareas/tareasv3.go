package main

import "fmt"

type taskList struct {
	// Slice de task
	tasks []*task
}

func (t *taskList) addTask(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) deleteTask(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

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
	t2 := &task{
		nombre:      "Curso Practico de Python",
		description: "Aprende Python con mini proyectos",
		completed:   false,
	}
	t3 := &task{
		nombre:      "Curso Practico de Java",
		description: "Aprende Java con mini proyectos",
		completed:   false,
	}
	lista := &taskList{
		tasks: []*task{
			t, t2,
		},
	}

	fmt.Println(lista.tasks[0])
	fmt.Println(lista.tasks[1])
	fmt.Println(len(lista.tasks))
	lista.addTask(t3)

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index", i, "Nombre", lista.tasks[i].nombre)
	}

	for index, tarea := range lista.tasks {
		fmt.Println("Index", index, "Nombre", tarea)
	}
	// tareas v2
	// fmt.Println(len(lista.tasks))
	// fmt.Println(lista.tasks[2])
	// lista.deleteTask(0)
	// fmt.Println(len(lista.tasks))
	// fmt.Println(lista.tasks[0])
	// fmt.Printf("%+v\n", t)

	// tareas v1
	// t.marcarCompleta()
	// t.newName("Curso Practico de Go Completo")
	// t.newDescription("Proyectos terminados")
	// fmt.Printf("%+v\n", t)

}
