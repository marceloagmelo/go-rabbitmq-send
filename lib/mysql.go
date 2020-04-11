package lib

import (
	"fmt"
	"log"
	"os"

	db "upper.io/db.v3"
	"upper.io/db.v3/mysql"
)

//Sess variavel que faz a conexão com o banco de dados
var Sess db.Database

func init() {
	var err error

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?allowNativePasswords=true", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOSTNAME"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	configuracao, err := mysql.ParseURL(connectionString)

	Sess, err = mysql.Open(configuracao)
	if err != nil {
		log.Fatal(err.Error())
	}
}
