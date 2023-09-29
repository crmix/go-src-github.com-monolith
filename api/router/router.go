package router

import (
	"monolith/api/controller"

	"github.com/gorilla/mux"
)

func Init(r *mux.Router, api controller.ProductAPI) {

	r.HandleFunc("/products", api.GetProductCtrl).Methods("GET")
}
