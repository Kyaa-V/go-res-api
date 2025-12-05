package service

import (
	"net/http"
	"fmt"
	"github.com/go-chi/chi/v5"
)

type Order struct {}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create order")
	 w.Write([]byte("Create Order service"))
}
func (o *Order) FindById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	 w.Write([]byte("Find order by id:" + id))
}	
func (o *Order) GetAll(w http.ResponseWriter, r *http.Request) {
	 w.Write([]byte("GetAll Order service"))
}
func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	 w.Write([]byte("update order by id:" + id))
}