package lib

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

const (
	fila string = "go-rabbitmq"
)

//ConectarRabbitMQ no rabbitmq
func ConectarRabbitMQ() (*amqp.Connection, error) {
	// Conectar com o rabbitmq
	var connectionString = fmt.Sprintf("amqp://%s:%s@%s:%s%s", os.Getenv("RABBITMQ_USER"), os.Getenv("RABBITMQ_PASS"), os.Getenv("RABBITMQ_HOSTNAME"), os.Getenv("RABBITMQ_PORT"), os.Getenv("RABBITMQ_VHOST"))
	conn, err := amqp.Dial(connectionString)
	if err != nil {
		log.Printf("ConectarRabbitMQ(): %s: %s", "Conectando com o rabbitmq", err)
		return nil, err
	}

	return conn, nil
}

//EnviarMensagemRabbitMQ no rabbitmq
func EnviarMensagemRabbitMQ(conn *amqp.Connection, novoID string) error {

	// Abrir o canal
	ch, err := conn.Channel()
	defer ch.Close()
	if err != nil {
		log.Printf("EnviarMensagemRabbitMQ(): %s: %s", "Abrindo canal no rabbitmq", err)
		return err
	}

	// Declarara fila
	q, err := ch.QueueDeclare(
		fila,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		log.Printf("EnviarMensagemRabbitMQ(): %s: %s", "Declarando fila", err)
		return err
	}

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
	if err != nil {
		log.Printf("EnviarMensagemRabbitMQ(): %s: %s", "Publicando mensagem", err)
		return err
	}

	log.Printf("Mensagem %s adicionada", body)

	return nil
}
