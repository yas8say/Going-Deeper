package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.Handle("/", recoverMiddleware(http.HandlerFunc(homeHandler)))
	http.Handle("/panic", recoverMiddleware(http.HandlerFunc(panicHandler)))

	http.ListenAndServe(":8080", nil)
}

func recoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Defer a function to recover from panic
		defer func() {
			if err := recover(); err != nil {
				// Return an internal server error response
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "Welocome to the homepage")
}

func panicHandler(w http.ResponseWriter, r *http.Request) {
	panic("something went wrong!")
}
