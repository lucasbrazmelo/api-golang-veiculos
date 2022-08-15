package veiculos

import "api/db"

func Get(id int64) (veiculo Veiculo, err error) {
	conn, err := db.Connect()
	if err != nil {
		return
	}
	defer conn.Close()

	row := conn.QueryRow(`SELECT * FROM veiculos WHERE idveiculos = ?`, id)

	err = row.Scan(
		&veiculo.Idveiculos,
		&veiculo.Nome,
		&veiculo.Marca,
		&veiculo.Ano,
		&veiculo.ValorVenda,
	)

	return
}