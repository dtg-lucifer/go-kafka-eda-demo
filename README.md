# Go Kafka Demo: Event-Driven Architecture with Go and Kafka

This repository demonstrates how to build a simple event-driven application
using Go and Kafka. It serves as a practical example for developers looking to
implement event-driven architecture in their Go applications.

## Overview

This project showcases a basic implementation of event-driven architecture
where:

- A REST API receives comment submissions
- Comments are pushed to Kafka topics as events
- A separate consumer service processes these events asynchronously

## Architecture

```
┌─────────────┐     ┌─────────────┐     ┌─────────────┐
│  API Server │────▶│    Kafka    │────▶│  Consumer   │
│  (Go Fiber) │     │   Broker    │     │   Worker    │
└─────────────┘     └─────────────┘     └─────────────┘
```

The application follows these principles:

- **Decoupling**: The API server and consumer worker are completely separate
- **Asynchronous Processing**: Comments are processed asynchronously
- **Scalability**: Each component can be scaled independently

## Technologies Used

- **Go**: Core programming language
- **Fiber**: Fast, lightweight web framework for Go
- **Kafka**: Distributed event streaming platform
- **Sarama**: Go client library for Apache Kafka
- **Docker & Docker Compose**: For containerization and orchestration
- **Air**: Live reload for Go applications during development

## Getting Started

### Prerequisites

- Go 1.24+
- Docker and Docker Compose
- Air (for development)

### Running the Application

1. **Start Kafka and Zookeeper**:
   ```bash
   docker-compose up -d
   ```

2. **Run the API server**:
   ```bash
   make run
   # or for development with hot reload
   make dev
   ```

3. **Run the consumer worker**:
   ```bash
   go run consumer/worker.go
   ```

## Project Structure

- `main.go` - Application entry point
- `server.go` - Server configuration and setup
- `handlers/` - HTTP request handlers
- `models/` - Data models
- `producer/` - Kafka producer implementation
- `consumer/` - Kafka consumer implementation
- `docker-compose.yaml` - Docker configuration for Kafka and Zookeeper

## API Endpoints

- `GET /api/v1/health` - Health check endpoint
- `GET /api/v1/metrics` - Metrics dashboard
- `POST /api/v1/comments` - Create a new comment (which gets published to Kafka)

### Creating a Comment

```bash
curl -X POST http://localhost:8080/api/v1/comments \
  -H "Content-Type: application/json" \
  -d '{"author":"John Doe","content":"This is a sample comment"}'
```

## How It Works

1. The client submits a comment via the REST API
2. The API server assigns an ID to the comment
3. The comment is serialized to JSON and published to the Kafka "comments" topic
4. The consumer worker receives the event from Kafka and processes it
5. The API responds to the client with the created comment details

## Event-Driven Architecture Benefits

- **Loose Coupling**: Services communicate through events, not direct calls
- **Scalability**: Each component can scale independently
- **Resilience**: Temporary service outages don't affect the entire system
- **Extensibility**: New consumers can be added without modifying existing code

## Development

For local development with hot reload:

```bash
make dev
```

## Future Improvements

- Add authentication and authorization
- Implement more advanced error handling and retries
- Add database persistence for comments
- Create a web UI for viewing comments
- Add unit and integration tests

## License

[MIT License](LICENSE)
