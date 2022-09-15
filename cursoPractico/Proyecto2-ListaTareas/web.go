package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct {
}

func (escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	// modulo http
	answer, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(answer)
	// fmt.Println(answer.Body)

	e := escritorWeb{}
	// func copy - nativa de go
	io.Copy(e, answer.Body)
}
