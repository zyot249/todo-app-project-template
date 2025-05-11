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
├── cmd/                    # Application entry points
├── internal/               # Private application code
│   ├── app/                # Application layer (use cases)
│   ├── domain/             # Domain layer (business rules)
│   ├── ports/              # Ports layer (interfaces used to communicate with the application)
│   └── adapters/           # Adapters layer (methods to communicate with the outside world)
├── pkg/                    # Public libraries
├── config/                 # Configuration files
├── docs/                   # Documentation
└── tests/                  # Integration tests
```

## Getting Started

### Prerequisites

- Go 1.24 or later
- PostgreSQL
- Make (optional, for using Makefile commands)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/todo-app.git
   cd todo-app
   ```

2. Install dependencies:

   ```bash
   go mod download
   ```

3. Set up environment variables:

   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. Run the application:

   ```bash
   go run cmd/main.go
   ```

## Development

- Run tests:

  ```bash
  go test ./...
  ```

## API Documentation

Once the service is running, you can access the API documentation at:

```text
http://localhost:8080/api/{service}/doc/index.html
```

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
