package main

import (
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page!")
}

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to API endpoint!")
}

func HandleFacu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Facu server!")
}
