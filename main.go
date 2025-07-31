package main

import (
	"net/http"
	"url_shortener/handlers"
)

func main() {
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	println("Running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
