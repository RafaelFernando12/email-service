package main

import (
	"context"
	"email-service/domain/email"
	"email-service/internal/client/rabbitmq"
	"email-service/internal/http"
	"email-service/pkg/env"
	"email-service/pkg/log"
	"os"
	"os/signal"
)

const (
	applicationName = "email-service"
)

func main() {
	logger := log.NewLogger(applicationName, env.GetEnv(env.EnvLoggerLevel, env.DefaultLoggerLevel))

	err := env.CheckRequiredEnv(
		env.EnvRabbitmqServer,
		env.EnvRabbitmqPort,
		env.EnvRabbitmqUser,
		env.EnvRabbitmqPassword,
		env.EnvRabbitmqQueueName,
		env.EnvSmtpHost,
		env.EnvSmtpPort,
		env.EnvSmtpUser,
		env.EnvSmtpPassword,
	)

	if err != nil {
		logger.Fatal().Print(err)

		return
	}

	client := rabbitmq.NewRabbitMQClient(
		env.GetEnv(env.EnvRabbitmqServer),
		env.GetEnv(env.EnvRabbitmqPort),
		env.GetEnv(env.EnvRabbitmqUser),
		env.GetEnv(env.EnvRabbitmqPassword))

	emailService := email.NewEmailService(client, env.GetEnv(env.EnvSmtpHost), env.GetEnv(env.EnvSmtpPort), env.GetEnv(env.EnvSmtpUser), env.GetEnv(env.EnvSmtpPassword))
	go emailService.StartListener(context.TODO(), env.GetEnv(env.EnvRabbitmqQueueName))

	server := http.NewServer(env.GetEnv(env.EnvPort, env.DefaultPort), nil, logger)
	server.Start()

	<-interrupt()

	server.Shutdown()
}

func interrupt() chan os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	return c
}
