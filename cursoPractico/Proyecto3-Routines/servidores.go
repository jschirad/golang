package main

import (
	"fmt"
	"net/http"
)

func main() {
	servers := []string{
		"http://platzi.com",
		"http://google.com",
		"http://youtube.com",
		"http://instagram.com",
	}
	for _, server := range servers {
		testServer(server)
	}
}

func testServer(server string) {
	_, err := http.Get(server)
	if err != nil {
		fmt.Println("Servidor no disponible")
	} else {
		fmt.Println("Servidor 200 OK", server)
	}
}
