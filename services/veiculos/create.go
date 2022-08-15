package veiculos

import (
	veiculosModel "api/models/veiculos";
	"encoding/json";
	"net/http";
	"fmt";
	"log";
)

func Create(res http.ResponseWriter, req *http.Request) {
	var veiculo veiculosModel.Veiculo

	err := json.NewDecoder(req.Body).Decode(&veiculo)
	if err != nil{
		log.Printf("Error decode do json: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := veiculosModel.Insert(veiculo)

	var resp map[string]any
	if	err != nil {
		resp = map[string]any{"messagem": fmt.Sprintf("Ocorreu um erro eu inserir %v  %s", err, id)}
	}else{
		resp = map[string]any{"error": false,"messagem": fmt.Sprintf("Veiculo inserido com sucesso %v %s", err, id)}
	}

	res.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(res).Encode(resp)
}