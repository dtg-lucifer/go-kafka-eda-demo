package handlers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/dtg-lucifer/go-kafka-demo/models"
	"github.com/dtg-lucifer/go-kafka-demo/producer"
)

func CreateComment(c *fiber.Ctx) error {
	comment := new(models.Comment)
	err := c.BodyParser(&comment)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not generate ID",
		})
	}
	comment.ID = id.String()

	commentInBytes, err := json.Marshal(comment)
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not marshal comment",
		})
	}
	log.Printf("Comment created: %s", string(commentInBytes))

	// @INFO - Push the comment to the queue
	err = producer.PushToQueue(producer.ProducerDTO{
		Topic: "comments",
		Data:  commentInBytes,
	})
	if err != nil {
		log.Print(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(comment)
}
