package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/dtg-lucifer/go-kafka-demo/handlers"
)

// ServerConfig holds the configuration for the Fiber server.
// It includes the port, host, and API version.
type ServerConfig struct {
	Port       int
	Host       string
	ApiVersion string
}

// Server represents the Fiber server configuration.
// It contains the port, host, API version, and the Fiber app instance.
// The Router is used to define the API routes.
type Server struct {
	Port       int
	Host       string
	ApiVersion string
	App        *fiber.App
	Router     fiber.Router
}

// NewServer initializes a new Fiber server with the given configuration.
// It creates a new Fiber app instance and sets up the router with the specified API version.
// The server is configured with the provided port and host.
func NewServer(cfg ServerConfig) *Server {
	app := fiber.New()
	router := app.Group(cfg.ApiVersion)

	return &Server{
		Port:   cfg.Port,
		Host:   cfg.Host,
		App:    app,
		Router: router,
	}
}

// Start initializes the server and starts listening on the specified host and port.
// It uses the Listen method of the Fiber app to start the server.
func (s *Server) Start() error {
	return s.App.Listen(s.Host + ":" + strconv.Itoa(s.Port))
}

// Init initializes the server by setting up middleware and routes.
// It uses the Fiber middleware for error handling, request ID generation, security headers,
// logging, and metrics.
func (s *Server) Init() {
	// @INFO - Middleware
	s.App.Use(recover.New())
	s.App.Use(requestid.New())
	s.App.Use(helmet.New())
	s.App.Use(logger.New(logger.Config{
		Format: "[${ip}]:${locals:requestid} - ${method} ${status} ${path} ${time} - ${latency} - ${error}\n",
	}))

	// @INFO - Routes
	s.SetupRoutes()

}

func (s *Server) SetupRoutes() {
	// @INFO - Metrics & Healthcheck
	s.Router.Get("/metrics", monitor.New())
	s.Router.Get("/health", handlers.HealthCheck)

	// @INFO - Comments
	s.Router.Post("/comments", handlers.CreateComment)
}
