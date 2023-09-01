package rabbitmq

import (
	"fmt"
	queries "shared-library/database/queries/queries-meme-uploader"

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
		queries.InsertMemePhotoId(string(d.Body))

	}

	return msgs
}
