package services

import (
	_ "api/services/veiculos"
	"net/http"
)

func Index(res http.ResponseWriter, req *http.Request) {

	http.ServeFile(res, req, "./views/veiculos/index.html")
	
}