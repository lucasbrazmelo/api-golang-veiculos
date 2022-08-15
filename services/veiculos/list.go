package veiculos

import (
	veiculosModel "api/models/veiculos";
	"encoding/json";
	"net/http";
	"log";
)

func List(res http.ResponseWriter, req *http.Request) {
	veiculos, err := veiculosModel.GetAll()
	if err != nil{
		log.Printf("Error obter veiculos: %v", err)
	}

	res.Header().Add("Content-Type", "application/josn")
	json.NewEncoder(res).Encode(veiculos)
}