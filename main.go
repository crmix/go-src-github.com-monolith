package main

import (
	"log"
	"monolith/api/controller"
	"monolith/api/router"
	"monolith/config"
	"monolith/platform"
	"monolith/service"
	"monolith/storage"
	"net/http"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

func main() {
	db := platform.DBConn()
	cfg := config.Load()

	storage := storage.NewProductRepo(db)
	service := service.NewProductService(storage, *zap.NewStdLog(zap.L()))
	api := controller.NewProductAPI(service)

	root := mux.NewRouter()

	router.Init(root, api)

	httpServer := http.Server{
		Addr:    cfg.HTTPPort,
		Handler: root,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("shut down server")
	}
}
