package veiculos

import (
	veiculosModel "api/models/veiculos";
	"encoding/json";
	"net/http";
	"log";
	"strconv";
	"github.com/go-chi/chi/v5";
)

func Get(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil{
		log.Printf("Error parse do id: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	veiculo, err := veiculosModel.Get(int64(id))
	if err != nil{
		log.Printf("Error retornar registro: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	res.Header().Add("Content-Type", "application/josn")
	json.NewEncoder(res).Encode(veiculo)
}