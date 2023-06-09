package database

import (
	"api_echo_modelo/src/configs"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Conectar() (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", configs.StringConexao)
	if err != nil {
		return nil, err
	}

	return db, nil
}
