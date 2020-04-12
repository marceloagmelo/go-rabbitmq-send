package models

import (
	"fmt"
	"log"

	"github.com/marceloagmelo/go-rabbitmq-send/lib"
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
	Criar(mensagemModel db.Collection) error
}

//Criar uma mensagem no banco de dados
func (m Mensagem) Criar(mensagemModel db.Collection) error {
	novoID, err := mensagemModel.Insert(m)
	if err != nil {
		log.Printf("Criar(): %s: %s", "Gravando a mensagem no banco de dados", err)
		return err
	}
	strID := fmt.Sprintf("%v", novoID)
	conn, err := lib.ConectarRabbitMQ()
	if err != nil {
		return err
	}
	defer conn.Close()

	err = lib.EnviarMensagemRabbitMQ(conn, strID)
	if err != nil {
		return err
	}
	return nil
}
