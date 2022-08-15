package veiculos

import (
	veiculosModel "api/models/veiculos";
	"encoding/json";
	"net/http";
	"fmt";
	"log";
	"strconv";
	"github.com/go-chi/chi/v5";
)

func Delete(res http.ResponseWriter, req *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(req, "id"))
	if err != nil{
		log.Printf("Error parse do id: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := veiculosModel.Delete(int64(id))
	if err != nil{
		log.Printf("Error apagar registro: %v", err)
		http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if	rows > 1 {
		log.Printf("Error apagar registro +1: %v", err)
	}
	
	resp := map[string]any{
			"error": false,
			"messagem": fmt.Sprintf("Veiculo apagado com sucesso %v", err),
	}


	res.Header().Add("Content-Type", "application/josn")
	json.NewEncoder(res).Encode(resp)
}