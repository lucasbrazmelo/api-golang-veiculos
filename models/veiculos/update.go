package veiculos

import "api/db"

func Update(id int64, veiculo Veiculo) (int64, error) {
	conn, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	res, err := conn.Exec(`UPDATE veiculos SET nome=?, marca=?, ano=?, valorvenda=? WHERE idveiculos=?`, veiculo.Nome, veiculo.Marca, veiculo.Ano, veiculo.ValorVenda, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}