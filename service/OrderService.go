package service

import (
	"net/http"
)

type Order struct {}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	 w.Write([]byte("Create Order service"))
}
func (o *Order) FindById(w http.ResponseWriter, r *http.Request) {
	 w.Write([]byte("FindById Order service"))
}
func (o *Order) GetAll(w http.ResponseWriter, r *http.Request) {
	 w.Write([]byte("GetAll Order service"))
}