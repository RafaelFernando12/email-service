# Golang Email Sender com RabbitMQ

Este repositório contém um serviço em **Golang** com o objetivo de consumir mensagens de uma fila no **RabbitMQ** e enviá-las por e-mail. Este projeto é parte de uma prática sobre sistemas de mensageria e funciona em conjunto com o serviço **SpringNoSql** (Java), cujo repositório pode ser encontrado na mesma organização no GitHub.

## Finalidade
- Praticar conceitos de **mensageria** usando RabbitMQ;
- Implementar um serviço simples de envio de e-mails com base nas mensagens recebidas de uma fila.

---

## Funcionalidades
- Conexão com uma fila RabbitMQ para consumo de mensagens;
- Envio de e-mails usando um servidor SMTP configurável;
- Configuração via variáveis de ambiente.

---

## Tecnologias Utilizadas
- **Golang**
- **RabbitMQ** (mensageria)
- **SMTP** (para envio de e-mails)

---

## Variáveis de Ambiente
Para executar o serviço, defina as seguintes variáveis de ambiente:

```env
# Configuração do servidor SMTP
SMTP_SERVER=smtp.gmail.com
SMTP_PORT=587
SMTP_USER=seu_email
SMTP_PASSWORD=sua_senha

# Configuração do servidor RabbitMQ
RABBITMQ_SERVER=localhost
RABBITMQ_PORT=5672
RABBITMQ_USER=guest
RABBITMQ_PASSWORD=guest
RABBITMQ_QUEUE_NAME=person_notifications
```

---

## Instalação e Execução

1. **Clone o repositório**:
   ```bash
   git clone https://github.com/RafaelFernando12/email-service.git
   cd email-service
   ```

2. **Configure as variáveis de ambiente**:
   Crie um arquivo `.env` ou exporte as variáveis no terminal conforme indicado acima.

3. **Execute o serviço**:
   ```bash
   go run main.go
   ```

4. **Testando**:
   - Garanta que o RabbitMQ esteja rodando na máquina configurada.
   - Publique mensagens na fila `person_notifications`.
   - Verifique o recebimento dos e-mails no servidor SMTP configurado.

---

## Estrutura do Projeto
```
├── main.go                # Ponto de entrada do serviço
  ├── domain/              # Lógica interna: consumo da mensagem e envio do e-mail
│   ├── email/             # Camada de implementação da lógica do sistema
│   ├── email.go           # Lógica de envio de e-mails
│   └── ...
├── internal/              # Pacotes internos do projeto
│   ├── client/            # Configuração e inicialização do cliente RabbitMQ
│   └── ...
└── go.mod                 # Gerenciamento de dependências
```

---

## Requisitos
- **Golang** instalado (v1.20+);
- **RabbitMQ** configurado e em execução;
- Um servidor SMTP funcional (ex.: Gmail, Outlook).

---

## Observações
- Para utilizar o **Gmail** como servidor SMTP, é necessário permitir o uso de "Aplicativos menos seguros" na conta.
- Caso deseje testar localmente sem um servidor SMTP, utilize ferramentas como [MailHog](https://github.com/mailhog/MailHog) ou [SMTP4Dev](https://github.com/rnwood/smtp4dev).

---

## Serviço Relacionado
- [**SpringNoSql**](https://github.com/RafaelFernando12/SpringNoSql) - Serviço em Java que complementa este projeto.
---

