package rabbitmq

import (
	"fmt"
	queries "shared-library/database/queries/queries-meme"

	"github.com/streadway/amqp"
)

func SendMessage(message, queueName string) {
	rabbitMQConn := connectToRabbitMQ(amqpURI)
	defer rabbitMQConn.conn.Close()

	rabbitMQChannel := createChannel(rabbitMQConn)
	defer rabbitMQChannel.ch.Close()

	err := rabbitMQChannel.ch.Publish(
		"",
		queueName,
		false,
		false,
		// Message type and content
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	failOnError(err, "Failed to publish message ")
}

func ReceivePhotoId(queue string) <-chan amqp.Delivery {
	rabbitMQConn := connectToRabbitMQ(amqpURI)
	defer rabbitMQConn.conn.Close()

	rabbitMQChannel := createChannel(rabbitMQConn)
	defer rabbitMQChannel.ch.Close()

	msgs, err := rabbitMQChannel.ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to consume messages")

	for d := range msgs {
		fmt.Printf("Received: %s\n", d.Body)
		err := queries.PhotoIdInsert(string(d.Body))
		if err != nil {
			d.Acknowledger.Nack(0, false, false)
		}
		d.Acknowledger.Ack(1, true)
	}

	return msgs
}

func ReceiveText(queue string) <-chan amqp.Delivery {
	rabbitMQConn := connectToRabbitMQ(amqpURI)
	defer rabbitMQConn.conn.Close()

	rabbitMQChannel := createChannel(rabbitMQConn)
	defer rabbitMQChannel.ch.Close()

	msgs, err := rabbitMQChannel.ch.Consume(
		queue,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to consume messages")

	for d := range msgs {
		fmt.Printf("Received: %s\n", d.Body)

		err := queries.TextInsert(string(d.Body))
		if err != nil {
			d.Acknowledger.Nack(0, false, false)
		}

		fmt.Println("Success textq: ", string(d.Body))
		d.Acknowledger.Ack(1, true)
	}

	return msgs
}
