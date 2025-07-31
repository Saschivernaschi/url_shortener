package handlers

import (
	"fmt"
	"net/http"
	"url_shortener/storage"

	"github.com/wordgen/wordgen"
)

func generateCode(n int) (string, error) {
	generator := wordgen.NewGenerator()
	var url_code string
	for i := 0; i < n; {
		word, err := generator.Generate()
		if err != nil {
			return "", err
		}
		url_code = url_code + "-" + word
	}
	return url_code, nil
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	url := r.FormValue("url")
	if url == "" {
		http.Error(w, "No URL provided", http.StatusBadRequest)
		return
	}
	code, err := generateCode(3)
	if err != nil {
		http.Error(w, "Could not generate short url", http.StatusInternalServerError)
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
