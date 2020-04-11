package lib

import (
	"log"

	db "upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

var configuracao = mysql.ConnectionURL{
	Host:     "localhost",
	User:     "root",
	Password: "admin123",
	Database: "membros",
}

//Sess variavel que faz a conexão com o banco de dados
var Sess db.Database

func init() {
	var err error
	Sess, err = mysql.Open(configuracao)
	if err != nil {
		log.Fatal(err.Error())
	}
}
