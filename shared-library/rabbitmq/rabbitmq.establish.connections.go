package rabbitmq

import "github.com/streadway/amqp"

type RabbitMQChannel struct {
	ch *amqp.Channel
}

type RabbitMQQueue struct {
	q amqp.Queue
}

func createChannel(rabbitMQConn RabbitMQConnection) RabbitMQChannel {
	ch, err := rabbitMQConn.conn.Channel()
	failOnError(err, "Failed to create channel ")
	return RabbitMQChannel{ch: ch}
}

func declareQueue(rabbitMQChannel RabbitMQChannel, queueName string) RabbitMQQueue {
	q, err := rabbitMQChannel.ch.QueueDeclare(
		queueName,
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare queue")
	return RabbitMQQueue{q: q}
}
