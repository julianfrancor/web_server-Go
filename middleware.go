package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// CheckAuth Este va a ser nuestro primer middleware, cadena de middlewares
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// con esta flag estamos simulando que la autenticacion fue correcta
			flag := true
			fmt.Println("Checking Authentication")
			if flag {
				f(w, r)
			} else {
				return
			}
		}
	}
}

// Logging which request was called, in which direction and at what time
func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
				// the parenthesis dow here is because that is an anonymous fucntion
			}()
			f(w, r)
		}
	}
}
