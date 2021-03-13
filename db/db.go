package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver
)

//Conectar abre a conexão com o banco de dados
func Conectar() *sql.DB {
	conexao := ":@/alura?charset=utf8&parseTime=True&loc=Local"
	//"user=golang dbname=alura password=golang host=localhosts sslmode=disable"
	db, erro := sql.Open("mysql", conexao)
	if erro != nil {
		panic(erro.Error())
	}
	return db
}
