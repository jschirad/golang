package main

import (
	"fmt"
	"net/http"
)

func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Acceso al escritor y al request -
		return func(w http.ResponseWriter, r *http.Request) {
			// Logica del Middleware
			flag := true
			fmt.Println("Checking Authentication")
			if flag {
				// Siguiente Middleware
				f(w, r)
			} else {
				return
			}
		}
	}
}
