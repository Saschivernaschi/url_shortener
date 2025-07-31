package handlers

import (
	"fmt"
	"math/rand"
	"net/http"
	"url_shortener/storage"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateCode(n int) string {
	code := make([]byte, n)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "No URL provided", http.StatusBadRequest)
		return
	}
	code := generateCode(6)
	storage.Save(code, url)
	fmt.Fprintf(w, "Short URL: http://localhost:8080/%s\n", code)
}

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:]
	if url, ok := storage.Get(code); ok {
		http.Redirect(w, r, url, http.StatusFound)
	} else {
		http.NotFound(w, r)
	}
}
