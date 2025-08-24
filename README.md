# Pseudo Links - Link Shortening Service

A link shortening service built with Go, featuring monitoring, logging, and containerization.

## Project Overview

Pseudo Links is a link shortening service that allows users to create short URLs that redirect to longer ones. The service includes comprehensive monitoring and logging capabilities using industry-standard tools.

## Technology Stack

### Backend

- **Go** - Main programming language
- **Gin Framework** - HTTP web framework
- **Redis** - In-memory data structure store for link storage
- **Docker** - Containerization platform

### Monitoring & Logging

- **Prometheus** - Metrics collection and monitoring
- **Grafana** - Data visualization and dashboard
- **Loki** - Log aggregation system

### API Documentation

- **Swagger/OpenAPI** - API documentation [`http://localhost:8080/swagger/index.html`]

## Project Structure

```
backend/
├── main-service/           # Go application
│   ├── internal/           # Core application logic
│   │   ├── link/           # Link management
│   │   ├── database/       # Database connections
│   │   ├── logger/         # Logging functionality
│   │   └── metrics/        # Metrics collection
│   ├── docs/               # Swagger documentation
│   └── main.go             # Application entry point
├── docker-compose.yml      # Service orchestration
├── prometheus.yml          # Prometheus configuration
├── loki-config.yml         # Loki configuration
└── grafana-datasources.yml # Grafana configuration
```

## Features

- Create short URLs from long ones
- Retrieve original URLs from short codes
- Delete existing links
- RESTful API with Swagger documentation
- Comprehensive logging with Loki
- Metrics monitoring with Prometheus
- Data visualization with Grafana dashboards
