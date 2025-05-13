# Todo Web Service

A RESTful web service built with Go, designed to manage todo items. This project serves as both a learning exercise for building web services in Go and a template for future web service development.

## Features

- RESTful API endpoints for CRUD operations on todo items
- Clean architecture pattern for maintainable and testable code
- Database integration for persistent storage
- Authentication and authorization
- API documentation
- Comprehensive test coverage

## Technical Stack

- **Language**: Go
- **Framework**: Standard library with custom middleware
- **Database**: PostgreSQL
- **Testing**: Go's testing package
- **Documentation**: Swagger/OpenAPI

## Project Structure

```text
todo-app/
├── docker/                 # Docker files for deployment
├── api/                    # Public APIs that the project exposes
├── cmd/                    # Application entry points
├── internal/               # Private application code
│   ├── {service}/          # Service code
│   │   ├── app/            # Application layer (use cases)
│   │   ├── domain/         # Domain layer (business rules)
│   │   ├── ports/          # Ports layer (interfaces used to communicate with the application)
│   │   └── adapters/       # Adapters layer (methods to communicate with the outside world)
├── pkg/                    # Public libraries
├── tool/                   # Go tools
├── config/                 # Configuration files
├── docs/                   # Guides and documentation
├── scripts/                # Scripts for building and running the project that are used in Makefile
└── tests/                  # Integration tests
```

## Getting Started

### Prerequisites

- Go 1.24 or later
- Docker
- Make (optional, for using Makefile commands)

### Installation

## Development

## API Documentation

Once the service is running, you can access the API documentation at:

```text
http://localhost:8080/api/{service}/doc/index.html
```

You can also use the swaggerui to test the api.
Swaggerui contains all the api documentation of services in the project.

```text
http://localhost:8000/doc/
```

## License

This project is licensed under the MIT License - see the LICENSE file for details.
