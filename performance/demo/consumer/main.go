package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"github.com/miloskrca/golang-training/performance/demo/app/rabbitmq"
	"github.com/streadway/amqp"
)

var (
	numQueues int
)

type RabbitMQConf struct {
	VHost    string   `yaml:"vhost"`
	User     string   `yaml:"user"`
	Password string   `yaml:"password"`
	Hostname string   `yaml:"hostname"`
	Port     string   `yaml:"port"`
	Exchange string   `yaml:"exchange"`
	Queues   []string `yaml:"queues"`
}

func main() {
	flag.IntVar(&numQueues, "q", 10, "num of queues")
	flag.Parse()

	var queueNames []string
	for i := 0; i < numQueues; i++ {
		queueNames = append(queueNames, fmt.Sprintf("client_%d", i))
	}
	config := RabbitMQConf{
		Hostname: "localhost",
		Port:     "5672",
		User:     "guest",
		Password: "guest",
		VHost:    "performance",
		Queues:   queueNames,
		Exchange: "performance_exchange",
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/%s", config.User, config.Password, config.Hostname, config.Port, config.VHost))
	if err != nil {
		log.Fatal(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	consumers, err := createConsumers(ch, queueNames)
	if err != nil {
		log.Fatal("could not create consumers", err)
		return
	}

	status := make(map[string]int, len(consumers))
	msgReceivedCh := make(chan string)
	for queue, consumer := range consumers {
		go runConsumer(queue, consumer, msgReceivedCh)
	}

	printStatus(status)
	for msgReceivedForConsumer := range msgReceivedCh {
		consumerCnt, found := status[msgReceivedForConsumer]
		if !found {
			status[msgReceivedForConsumer] = 0
		}
		status[msgReceivedForConsumer] = consumerCnt + 1
		printStatus(status)
	}
}

func runConsumer(queue string, consumer <-chan amqp.Delivery, msgReceivedCh chan string) {
	for consumedMessage := range consumer {
		var message rabbitmq.Message
		err := json.Unmarshal(consumedMessage.Body, &message)
		if err != nil {
			log.Printf("could not unmarshal message: %v", err)
			continue
		}
		msgReceivedCh <- queue
	}
}

func createConsumers(ch *amqp.Channel, queues []string) (map[string]<-chan amqp.Delivery, error) {
	consumers := make(map[string]<-chan amqp.Delivery)
	for _, queueName := range queues {
		consumer, err := ch.Consume(queueName, "", true, false, false, false, nil)
		if err != nil {
			return nil, err
		}
		consumers[queueName] = consumer
	}
	return consumers, nil
}

func printStatus(status map[string]int) {
	var total int
	for _, count := range status {
		total = total + count
	}
	fmt.Printf("\033[2K\rTotal received messages: %d", total)
}
