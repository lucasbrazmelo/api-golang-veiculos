package db

import (
	"database/sql";
	"api/configs";
	"fmt";
	_ "github.com/go-sql-driver/mysql";
)

func Connect() (*sql.DB, error) {
	config := configs.GetDB()
	configString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", 
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
	)

	conn, err := sql.Open("mysql", configString)

	if err != nil {
		panic(err)
	}

	err = conn.Ping()

	return conn, err
}