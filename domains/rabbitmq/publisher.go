package rabbitmq

import (
	"restaurant-api/utils"

	"github.com/streadway/amqp"
)

func PublishMessage(message string, queueName string) error {
	err := utils.Channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return err
}
