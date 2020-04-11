package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/marceloagmelo/go-rabbitmq-send/utils"
	"github.com/streadway/amqp"
)

const (
	fila string = "go-rabbitmq"
)

//ConectarRabbitMQ no rabbitmq
func ConectarRabbitMQ() (conn *amqp.Connection, mensagem string) {
	// Conectar com o rabbitmq
	var connectionString = fmt.Sprintf("amqp://%s:%s@%s:%s%s", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), os.Getenv("RABBITMQ_HOSTNAME"), os.Getenv("RABBITMQ_PORT"), os.Getenv("RABBITMQ_VHOST"))
	conn, err := amqp.Dial(connectionString)
	mensagem = utils.CheckErr(err, "Conectando com o rabbitmq")

	if mensagem != "" {
		log.Println(mensagem)
	}

	return conn, mensagem
}

//EnviarMensagemRabbitMQ no rabbitmq
func EnviarMensagemRabbitMQ(conn *amqp.Connection, novoID string) string {
	mensagem := ""

	// Abrir o canal
	ch, err := conn.Channel()
	mensagem = utils.CheckErr(err, "Abrindo canal")

	defer ch.Close()

	// Declarara fila
	q, err := ch.QueueDeclare(
		"go-rabbitmq", // name
		false,         // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	mensagem = utils.CheckErr(err, "Declarando fila")

	body := novoID
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	mensagem = utils.CheckErr(err, "Publicando mensagem")

	if mensagem == "" {
		mensagem = fmt.Sprintf("Mensagem %s adicionada", body)
		log.Println(mensagem)
	}

	return mensagem
}
