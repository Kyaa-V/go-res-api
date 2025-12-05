package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/Kyaa-V/go-res-api/service"
)

func LoadOrderController(r chi.Router ) {

	orderService := &service.Order{}

	r.Post("/", orderService.Create)
}