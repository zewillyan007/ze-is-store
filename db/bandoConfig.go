package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConectionMySql() *sql.DB {
	//conexao := "zewillyan@willyan2014(localhost)/loja_ze?charset=utf8&parseTime=True&loc=Local"
	conexao := "zewillyan:willyan2014@tcp(localhost:3306)/loja_ze"
	db, err := sql.Open("mysql", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
