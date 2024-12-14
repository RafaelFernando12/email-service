package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQClient interface {
	Connect() (*amqp.Connection, error)
	Consume(queueName string) (<-chan amqp.Delivery, error)
}

type rabbitMQClient struct {
	server   string
	port     string
	user     string
	password string
}

func NewRabbitMQClient(server, port, user, password string) *rabbitMQClient {
	return &rabbitMQClient{
		server:   server,
		port:     port,
		user:     user,
		password: password,
	}
}

func (c *rabbitMQClient) Connect() (*amqp.Connection, error) {
	connectionURL := fmt.Sprintf("amqp://%s:%s@%s:%s/", c.user, c.password, c.server, c.port)
	conn, err := amqp.Dial(connectionURL)
	if err != nil {
		log.Printf("Erro ao conectar ao RabbitMQ: %v", err)
		return nil, err
	}

	log.Println("Conexão com RabbitMQ estabelecida com sucesso")
	return conn, nil
}

func (c *rabbitMQClient) Consume(queueName string) (<-chan amqp.Delivery, error) {
	conn, err := c.Connect()
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Printf("Erro ao criar canal RabbitMQ: %v", err)
		return nil, err
	}

	_, err = channel.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Erro ao declarar fila: %v", err)
		return nil, err
	}

	messages, err := channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("Erro ao iniciar consumo: %v", err)
		return nil, err
	}

	log.Printf("Consumindo mensagens da fila '%s'", queueName)
	return messages, nil
}
