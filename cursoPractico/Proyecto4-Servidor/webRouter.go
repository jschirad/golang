package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handle, exist := r.rules[path]
	return handle, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// fmt.Fprintf(w, "Hello World!")
	handler, exist := r.FindHandler(request.URL.Path)
	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}
