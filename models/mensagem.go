package models

import (
	"fmt"

	"github.com/marceloagmelo/go-rabbitmq-send/lib"
	"github.com/marceloagmelo/go-rabbitmq-send/utils"
)

//Mensagem estrutura de mensagem
type Mensagem struct {
	ID     int    `db:"id" json:"id"`
	Titulo string `db:"titulo" json:"titulo"`
	Texto  string `db:"texto" json:"texto"`
	/*DataCriacao     time.Time `db:"dtcriacao" json:"dtcriacao"`
	DataAtualizacao time.Time `db:"dtatualizacao" json:"dtatualizacao"`*/
	Status int `db:"status" json:"status"`
}

// Metodos interface
type Metodos interface {
	Criar() string
}

//MensagemModel recebe a tabela do banco de dados
var MensagemModel = lib.Sess.Collection("mensagem")

//Criar uma mensagem no banco de dados
func (m Mensagem) Criar() string {
	novoID, err := MensagemModel.Insert(m)
	mensagem := utils.CheckErr(err, "Conectando com o rabbitmq")
	if mensagem == "" {
		conn, mensagem := lib.ConectarRabbitMQ()
		defer conn.Close()
		if mensagem == "" {
			strID := fmt.Sprintf("%v", novoID)
			mensagem = lib.EnviarMensagemRabbitMQ(conn, strID)
		}
	}
	return mensagem
}
