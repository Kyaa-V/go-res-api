package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/Kyaa-V/go-res-api/service"
	"fmt"
)

func LoadOrderController(r chi.Router ) {

	orderService := &service.Order{}
	fmt.Println("load order controller routes")
	r.Post("/", orderService.Create)
	r.Get("/", orderService.GetAll)
	r.Get("/{id}", orderService.FindById)
	r.Put("/{id}", orderService.UpdateById)
}