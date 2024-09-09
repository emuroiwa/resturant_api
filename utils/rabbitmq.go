package utils

import (
	"log"

	"github.com/streadway/amqp"
)

var Channel *amqp.Channel
var RatingQueue amqp.Queue
var DishQueue amqp.Queue

// SetupRabbitMQ initializes the RabbitMQ connection and declares the queues
func SetupRabbitMQ() error {
	// Connect to RabbitMQ server
	conn, err := amqp.Dial("amqp://admin:admin123@localhost:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return err
	}

	// Create a channel
	Channel, err = conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return err
	}

	RatingQueue, err = Channel.QueueDeclare(
		"rating_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare rating queue: %v", err)
		return err
	}

	DishQueue, err = Channel.QueueDeclare(
		"dish_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to declare dish queue: %v", err)
		return err
	}

	log.Println("RabbitMQ setup complete. Queues: rating_queue, dish_queue")
	return nil
}
