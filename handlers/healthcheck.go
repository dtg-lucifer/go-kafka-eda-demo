package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

var StartTime = time.Now()

func HealthCheck(c *fiber.Ctx) error {
	requestId := uuid.New()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":    "UP",
		"message":   "Service is running",
		"uptime":    time.Since(StartTime).String(),
		"timestamp": time.Now().Format(time.RFC3339),
		"requestId": requestId.String(),
	})
}
