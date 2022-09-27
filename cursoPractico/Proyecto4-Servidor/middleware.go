package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Middleware Template
// func MiddlewareName() Middleware {
// 	return func(f http.HandlerFunc) http.HandlerFunc {
// 		// Acceso al escritor y al request -
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			// Logica del Middleware

// 		}
// 	}
// }

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

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		// Acceso al escritor y al request -
		return func(w http.ResponseWriter, r *http.Request) {
			// Logica del Middleware
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			// Siguiente Middleware
			f(w, r)
		}
	}
}
