package dishes

import (
	"encoding/json"
	"log"
	"restaurant-api/models"
	"restaurant-api/repositories"
	"restaurant-api/utils"

	"gorm.io/gorm"
)

func StartConsumer(db *gorm.DB) {
	msgs, err := utils.Channel.Consume(
		"dish_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}

	go func() {
		for d := range msgs {
			var rating models.Dish
			err := json.Unmarshal(d.Body, &rating)
			if err != nil {
				log.Printf("Error decoding JSON: %v", err)
				continue
			}

			dishRepo := repositories.NewDishRepository(db)
			err = dishRepo.Create(&rating)
			if err != nil {
				log.Printf("Error saving rating: %v", err)
				continue
			}
			log.Println("Rating saved successfully")
		}
	}()

	log.Println("RabbitMQ consumer started")
}
