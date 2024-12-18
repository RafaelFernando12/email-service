package email

import (
	"context"
	"email-service/domain"
	"email-service/internal/client/rabbitmq"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"gopkg.in/mail.v2"
)

type emailService struct {
	rabbitmqClient                             rabbitmq.RabbitMQClient
	SMTPHost, SMTPPort, SMTPUser, SMTPPassword string
}

func NewEmailService(rabbitmq rabbitmq.RabbitMQClient, SMTPHost, SMTPPort, SMTPUser, SMTPPassword string) *emailService {
	return &emailService{
		rabbitmqClient: rabbitmq,
		SMTPHost:       SMTPHost,
		SMTPPort:       SMTPPort,
		SMTPUser:       SMTPUser,
		SMTPPassword:   SMTPPassword,
	}
}

func (e *emailService) Send(ctx context.Context, person *domain.Person) bool {
	log.Printf("Enviando e-mail para: %s <%s>", person.Name, person.Email)

	message := mail.NewMessage()
	message.SetHeader("From", "sistemarafa256@gmail.com")
	message.SetHeader("To", person.Email)
	message.SetHeader("Subject", "Bem-vindo ao Nosso Serviço")
	message.SetBody("text/plain", fmt.Sprintf("Olá %s, bem-vindo ao nosso serviço!", person.Name))

	port, err := strconv.Atoi(e.SMTPPort)
	if err != nil {
		fmt.Printf("Erro ao converter string para int: %v\n", err)
		return false
	}

	dialer := mail.NewDialer(e.SMTPHost, port, e.SMTPUser, e.SMTPPassword)

	if err := dialer.DialAndSend(message); err != nil {
		log.Printf("Erro ao enviar e-mail para %s: %v", person.Email, err)
		return false
	}

	log.Printf("E-mail enviado com sucesso para %s", person.Email)
	return true
}

func (e *emailService) StartListener(ctx context.Context, queueName string) {
	messages, err := e.rabbitmqClient.Consume(queueName)
	if err != nil {
		log.Fatalf("Erro ao consumir mensagens da fila '%s': %v", queueName, err)
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Listener encerrado pelo contexto.")
			return
		case msg := <-messages:
			var person domain.Person

			err := json.Unmarshal(msg.Body, &person)
			if err != nil {
				log.Printf("Erro ao converter mensagem para Person: %v", err)
				_ = msg.Nack(false, true)
				continue
			}

			success := e.Send(ctx, &person)

			if success {
				err := msg.Ack(false)
				if err != nil {
					log.Printf("Erro ao confirmar a mensagem: %v", err)
				}
			} else {
				err := msg.Nack(false, true)
				if err != nil {
					log.Printf("Erro ao rejeitar a mensagem: %v", err)
				}
			}
		}
	}
}
