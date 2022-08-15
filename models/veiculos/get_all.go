package veiculos

import "api/db"

func GetAll() (veiculos []Veiculo, err error) {
	conn, err := db.Connect()
	if err != nil {
		return
	}
	defer conn.Close()

	rows, err := conn.Query(`SELECT * FROM veiculos`)
	if err != nil {
		return
	}

	for rows.Next(){
		var veiculo Veiculo
		
		err = rows.Scan(
			&veiculo.Idveiculos,
			&veiculo.Nome,
			&veiculo.Marca,
			&veiculo.Ano,
			&veiculo.ValorVenda,
		)
		if err != nil {
			continue
		}

		veiculos = append(veiculos, veiculo)
	}
	
	return
}