package controller

import (
	"fmt"
	"log"
	"monolith/service"
	"net/http"
	"strconv"
)

type productApi struct {
	service service.ProductService
}

func NewProductAPI(service service.ProductService) ProductAPI {
	return &productApi{
		service: service,
	}
}

type ProductAPI interface {
	GetProductCtrl(w http.ResponseWriter, r *http.Request)
}

func (api *productApi) GetProductCtrl(w http.ResponseWriter, r *http.Request) {
	// v := mux.Vars(r)
	// ID, err := strconv.Atoi(v["id"])
	// if err != nil {
	// 	fmt.Fprintln(w, "error in convert from string to int")
	// }
     q :=r.URL.Query()

	 page,err :=strconv.Atoi(q.Get("page"))
     if err !=nil {
		fmt.Print(err)
	 }
	 limit, err :=strconv.Atoi(q.Get("limit"))
	   if err!=nil {
		fmt.Print(err)
	   }

	res, err := api.service.GetProductCtrl(page, limit)
	if err != nil {
		log.Println(err)
		return
	}

	WriteSuccess(w, res)
}
