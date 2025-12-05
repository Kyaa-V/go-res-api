package application

import (
	"github.com/go-chi/chi/v5"
	"github.com/Kyaa-V/go-res-api/controller"
	"net/http"
	"fmt"
)

func setupRouter() http.Handler {
	r := chi.NewRouter()
	fmt.Println("setup router")

	r.Get("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello world"))
	})

	r.Route("/orders", controller.LoadOrderController)

	return r
}