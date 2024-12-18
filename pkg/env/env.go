package env

import (
	"fmt"
	"os"
)

const (
	EnvPort        = "PORT"
	EnvLoggerLevel = "LOGGER_LEVEL"
	DefaultPort    = "8082"

	EnvSmtpHost     = "SMTP_SERVER"
	EnvSmtpPort     = "SMTP_PORT"
	EnvSmtpUser     = "SMTP_USER"
	EnvSmtpPassword = "SMTP_PASSWORD"

	EnvRabbitmqServer    = "RABBITMQ_SERVER"
	EnvRabbitmqPort      = "RABBITMQ_PORT"
	EnvRabbitmqUser      = "RABBITMQ_USER"
	EnvRabbitmqPassword  = "RABBITMQ_PASSWORD"
	EnvRabbitmqQueueName = "RABBITMQ_QUEUE_NAME"

	DefaultLoggerLevel = "DEBUG"
)

func GetEnv(args ...string) string {
	value := os.Getenv(args[0])
	if value == "" && len(args) > 1 {
		return args[1]
	}

	return value
}

func CheckRequiredEnv(envs ...string) error {
	for _, env := range envs {
		value := os.Getenv(env)
		if value == "" {
			return &EnvRequiredError{
				Env: env,
			}
		}
	}

	return nil
}

type EnvRequiredError struct {
	Env string
}

func (err *EnvRequiredError) Error() string {
	return fmt.Sprintf("env %s is required", err.Env)
}
