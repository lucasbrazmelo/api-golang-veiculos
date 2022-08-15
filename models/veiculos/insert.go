package veiculos

import "api/db"

func Insert(veiculo Veiculo) (id int64, err error) {
	conn, err := db.Connect()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `INSERT INTO veiculos (nome, marca, ano, valorvenda) VALUES (?, ?, ?, ?)`

	res, err := conn.Exec(sql, veiculo.Nome, veiculo.Marca, veiculo.Ano, veiculo.ValorVenda)
	if err != nil {
        println("Exec err:", err.Error())
    }
	id, err = res.LastInsertId()
	if err != nil {
        return 0, err
    }

	return id, err
}