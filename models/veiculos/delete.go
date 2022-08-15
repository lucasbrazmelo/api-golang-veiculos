package veiculos

import "api/db"

func Delete(id int64) (rowsAffected int64, err error) {
	conn, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	sql := `DELETE FROM veiculos WHERE idveiculos = ?`

	res, err := conn.Exec(sql, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err = res.RowsAffected()

	return rowsAffected, err
}