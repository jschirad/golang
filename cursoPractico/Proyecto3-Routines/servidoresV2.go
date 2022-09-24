package main

import (
	"fmt"
	"net/http"
)

func main() {
	canal := make(chan string)
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://youtube.com",
		"http://instagram.com",
	}
	for _, server := range servers {
		go testServer(server, canal)
	}

	for i := 0; i < len(servers); i++ {
		fmt.Println(<-canal)
	}
}

func testServer(server string, canal chan string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println("Servidor no disponible")
		canal <- server + " Server not found"
	} else {
		fmt.Println("Servidor 200 OK", server)
		canal <- server + " Status 200 OK"
	}
}
