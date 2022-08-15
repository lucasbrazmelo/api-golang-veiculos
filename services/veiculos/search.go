package veiculos

import (
	veiculosModel "api/models/veiculos";
	"encoding/json";
	"net/http";
	"log";
	"github.com/go-chi/chi/v5";
)

func Search(res http.ResponseWriter, req *http.Request) {

	veiculo, err := veiculosModel.Search(chi.URLParam(req, "search"))
	if err != nil{
		log.Printf("Error retornar registro: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.Header().Add("Content-Type", "application/josn")
	json.NewEncoder(res).Encode(veiculo)
}