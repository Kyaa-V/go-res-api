package application

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func setupRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello world"))
	})

	return r
}