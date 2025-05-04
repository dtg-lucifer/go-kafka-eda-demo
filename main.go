package main

import (
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Info("Error loading .env file, using system environment variables")
	}

	server := NewServer(ServerConfig{
		Port:       8080,
		Host:       "0.0.0.0",
		ApiVersion: "/api/v1",
	})

	server.Init()
	server.Start()
}
