package domain

import (
	"context"
	"email-service/internal/client/rabbitmq"
)

type emailService interface {
	Send(ctx context.Context, rabbitmqClient *rabbitmq.RabbitMQClient)
}
