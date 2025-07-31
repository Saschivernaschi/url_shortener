package handlers

import (
	"fmt"
	"net/http"
	"url_shortener/storage"

	"github.com/wordgen/wordgen"
	"github.com/wordgen/wordlists"
)

func generateCode() (string, error) {
	generator := wordgen.NewGenerator()
	generator.Words = wordlists.EffLarge

	word, err := generator.Generate()
	if err != nil {
		return "", err
	}

	return word, nil
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "No URL provided", http.StatusBadRequest)
		return
	}
	code, err := generateCode()
	if err != nil {
		http.Error(w, "Could not generate words for short url", http.StatusInternalServerError)
		return
	}

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
