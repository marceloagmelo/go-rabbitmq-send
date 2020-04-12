package models

import (
	"fmt"

	"github.com/marceloagmelo/go-rabbitmq-send/lib"
	"github.com/marceloagmelo/go-rabbitmq-send/utils"
	"upper.io/db.v3"
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
	Criar(mensagemModel db.Collection) string
}

//Criar uma mensagem no banco de dados
func (m Mensagem) Criar(mensagemModel db.Collection) string {
	novoID, err := mensagemModel.Insert(m)
	mensagem := utils.CheckErr(err, "Gravando mensagem no banco de dados")
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
