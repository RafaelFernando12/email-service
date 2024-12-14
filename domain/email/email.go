package email

import (
	"context"
	"email-service/internal/client/rabbitmq"
	"fmt"
	"time"
)

type emailService struct {
	rabbitmqClient *rabbitmq.RabbitMQClient
}

func NewEmailService(rabbitmq rabbitmq.RabbitMQClient) *emailService {
	return &emailService{
		rabbitmqClient: &rabbitmq,
	}
}

func (e *emailService) Send(ctx context.Context, client rabbitmq.RabbitMQClient) {
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-time.After(time.Second):
				client.Connect()
				messages, err := client.Consume("person_notifications")
				if err != nil {
					return
				}

				for msg := range messages {
					fmt.Print(msg)
					return
				}
			case <-done:
				return
			}
		}
	}()

	time.Sleep(time.Minute)
	done <- true

}
