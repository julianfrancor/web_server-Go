package main

import (
	"net/http"
)

type Router struct {
	// mapa que pase de stringd a handlers
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(method string, path string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, exist, methodExist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exit, methodExist := r.FindHandler(request.URL.Path, request.Method)

	if exit != true {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
