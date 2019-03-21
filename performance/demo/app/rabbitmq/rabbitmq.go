package rabbitmq

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type Message struct {
	QueueName string
	Type      string
	Payload   []byte
}

type RabbitMQConf struct {
	VHost    string   `yaml:"vhost"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	Hostname string   `yaml:"hostname"`
	Port     string   `yaml:"port"`
	Exchange string   `yaml:"exchange"`
	Queues   []string `yaml:"queues"`
}

type RabbitMQ struct {
	conn   *amqp.Connection
	config RabbitMQConf
	send   chan Message
}

const (
	deadLetterExchange         = "indexing-mdp.dlq"
	deadLetterRoutingKeyFormat = "%s.dlq"
)

func New(config RabbitMQConf) (*RabbitMQ, error) {
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.User, config.Password, config.Hostname, config.Port, config.VHost))
	if err != nil {
		return nil, err
	}
	return &RabbitMQ{conn, config, make(chan Message)}, nil
}

func (service *RabbitMQ) Start() {
	ch, err := service.conn.Channel()
	if err != nil {
		log.Printf("could not open channel: %v", err)
		return
	}

	err = service.setupExchange(ch)
	if err != nil {
		log.Printf("could not declare exchange, opening the channel again: %v", err)
		ch, err = service.conn.Channel()
		if err != nil {
			log.Printf("could not open channel after failed exchange declaration: %v", err)
			return
		}
	}

	err = service.setupQueues(ch)
	if err != nil {
		log.Printf("could not setup queues: %v", err)
		return
	}

	for msg := range service.send {
		// log.Printf("sending message of type %s to %s: %s", msg.Type, msg.QueueName, string(msg.Payload))

		err = ch.Publish(
			service.config.Exchange, // exchange
			msg.QueueName,           // routing key
			false,                   // mandatory
			false,                   // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body:        msg.Payload,
			})
		if err != nil {
			log.Printf("could not publish: %v", err)
			continue
		}
	}
}

func (service *RabbitMQ) Send(msg Message) {
	service.send <- msg
}

func (service *RabbitMQ) Stop() {
	close(service.send)
}

func (service *RabbitMQ) setupExchange(ch *amqp.Channel) error {
	err := ch.ExchangeDeclare(service.config.Exchange, "topic", true, false, false, false, nil)
	if err != nil {
		return err
	}
	return nil
}

func (service *RabbitMQ) setupQueues(ch *amqp.Channel) error {
	for _, queueName := range service.config.Queues {
		_, err := ch.QueueDeclare(queueName, true, false, false, false, map[string]interface{}{
			"x-dead-letter-exchange":    deadLetterExchange,
			"x-dead-letter-routing-key": fmt.Sprintf(deadLetterRoutingKeyFormat, queueName),
		})
		if err != nil {
			return fmt.Errorf("could not declare queue: %v", err)
		}

		if err := ch.QueueBind(queueName, queueName, service.config.Exchange, false, nil); err != nil {
			log.Printf("could not bind queue to exchange: %v", err)
			return fmt.Errorf("could not bind queue: %v", err)
		}
	}
	return nil
}
