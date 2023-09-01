package types

import "github.com/streadway/amqp"

type RabbitMQConnection struct {
	conn *amqp.Connection
}

type RabbitMQChannel struct {
	ch *amqp.Channel
}

type RabbitMQQueue struct {
	q amqp.Queue
}
