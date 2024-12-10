package rabbitmq

"github.com/streadway/amqp"

type RabbitMQClient interface {
	Connect() (*amqp.Connection, error)
}

type rabbitMQClient struct {
	url      string
	server   string
	port     string
	user     string
	password string
}

func NewRabbitMQClient(url, server, port, user, password string) *rabbitMQClient {
	return &rabbitMQClient{
		url:      url,
		server:   server,
		port:     port,
		user:     user,
		password: password,
	}
}

func (c *rabbitMQClient) Connect() (*amqp.Connection, error) {
	var connectionURL string

	if c.url != "" {
		connectionURL = c.url
	} else {
		connectionURL = fmt.Sprintf("amqp://%s:%s@%s:%s/", c.user, c.password, c.server, c.port)
	}

	conn, err := amqp.Dial(connectionURL)
	if err != nil {
		log.Printf("Erro ao conectar ao RabbitMQ: %v", err)
		return nil, err
	}

	log.Println("Conexão com RabbitMQ estabelecida com sucesso")
	return conn, nil
}