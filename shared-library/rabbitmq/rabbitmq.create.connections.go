package rabbitmq

import (
	"log"

	"github.com/streadway/amqp"
)

type RabbitMQConnection struct {
	conn *amqp.Connection
}

var amqpURI string = "amqp://localhost:5672/"

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func connectToRabbitMQ(amqpURI string) RabbitMQConnection {
	conn, err := amqp.Dial(amqpURI)
	failOnError(err, "Failed to establish connection")
	return RabbitMQConnection{conn: conn}
}

func QueueRabbitStart() {
	rabbitMQConn := connectToRabbitMQ(amqpURI)
	defer rabbitMQConn.conn.Close()

	rabbitMQChannel := createChannel(rabbitMQConn)
	defer rabbitMQChannel.ch.Close()

	declareQueue(rabbitMQChannel, "photoq")
	declareQueue(rabbitMQChannel, "textq")
}
