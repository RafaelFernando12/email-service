package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Erro ao carregar .env: %v", err)
	}

	rabbitmqURL := os.Getenv("RABBITMQ_URL")
	smtpServer := os.Getenv("SMTP_SERVER")
	log.Println("RabbitMQ URL:", rabbitmqURL)
	log.Println("SMTP Server:", smtpServer)
}
