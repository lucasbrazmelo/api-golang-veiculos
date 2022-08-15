package main

import (
	services "api/services";
	veiculos "api/services/veiculos";
	"api/configs";
	"net/http";
	"fmt";
	"github.com/go-chi/chi/v5";
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	router := chi.NewRouter()
	router.Post("/veiculos", veiculos.Create)
	router.Put("/veiculos/{id}", veiculos.Update)
	router.Delete("/veiculos/{id}", veiculos.Delete)
	router.Get("/veiculos", veiculos.List)
	router.Get("/veiculos/search/{search}", veiculos.Search)
	router.Get("/veiculos/{id}", veiculos.Get)
	router.Get("/", services.Index)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}