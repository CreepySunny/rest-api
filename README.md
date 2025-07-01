# Calculator Backend API

A production-ready calculator REST API built with Go that provides basic arithmetic operations. This project implements a stateless HTTP+JSON API that conforms to the provided OpenAPI specification.

## Features

- **Addition**: Add two integers
- **Subtraction**: Subtract two integers  
- **Multiplication**: Multiply two integers
- **Division**: Divide two integers with zero-division protection
- **Sum**: Calculate the sum of an array of integers
- **Input Validation**: Comprehensive request validation and sanitization
- **Error Handling**: Detailed error feedback for invalid inputs
- **Structured Logging**: Request logging with status codes, IP addresses, and paths
- **CORS Support**: Cross-origin resource sharing for web applications

## Production Ready Features

This API implements several production-ready features:

### Input Validation
- Comprehensive validation and sanitization of all user inputs
- Protection against division by zero operations
- Type validation for numeric inputs
- Array validation for sum operations

### Error Handling
- Detailed error responses with appropriate HTTP status codes
- Clear error messages to help users correct invalid inputs
- Graceful handling of edge cases and malformed requests

### Structured Logging
- Request logging using Go's `log/slog` package
- Logs include: timestamp, method, path, status code, IP address, and response time
- Available in both JSON and text formats for different environments

### CORS Support
- Cross-Origin Resource Sharing enabled for web applications
- Configurable CORS policies for different environments

## Reference Implementation

An example version of this API is available at: https://calculator.dreamsofcode.io

## API Endpoints

### Base URL
```
http://localhost:3000
```

### 1. Add Two Numbers
**POST** `/add`

Adds two integers together.

**Request Body:**
```json
{
  "number1": 10,
  "number2": 5
}
```

**Response:**
```json
{
  "result": 15
}
```

### 2. Subtract Two Numbers
**POST** `/subtract`

Subtracts the second number from the first number.

**Request Body:**
```json
{
  "number1": 10,
  "number2": 3
}
```

**Response:**
```json
{
  "result": 7
}
```

### 3. Multiply Two Numbers
**POST** `/multiply`

Multiplies two integers together.

**Request Body:**
```json
{
  "number1": 4,
  "number2": 6
}
```

**Response:**
```json
{
  "result": 24
}
```

### 4. Divide Two Numbers
**POST** `/divide`

Divides the dividend by the divisor.

**Request Body:**
```json
{
  "dividend": 20,
  "divisor": 4
}
```

**Response:**
```json
{
  "result": 5
}
```

### 5. Sum Array of Numbers
**POST** `/sum`

Calculates the sum of all numbers in an array.

**Request Body:**
```json
[1, 2, 3, 4, 5]
```

**Response:**
```json
{
  "result": 15
}
```

## Getting Started

### Prerequisites

- Go 1.19 or higher
- Git

### Recommended Packages

This implementation uses the following Go packages:

- `net/http` - Standard library HTTP package for routing and server setup
- `encoding/json` - JSON encoding/decoding for request/response bodies
- `log/slog` - Structured logging (JSON and text formats)
- `github.com/rs/cors` - CORS middleware for cross-origin requests

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd rest-api
```

2. Install dependencies:
```bash
go mod init calculator-api
go get github.com/rs/cors
go mod tidy
```

3. Run the server:
```bash
go run main.go
```

The server will start on `http://localhost:3000`.

### Configuration

The API supports configuration through environment variables:

- `PORT` - Server port (default: 3000)
- `LOG_FORMAT` - Logging format: "json" or "text" (default: "text")
- `CORS_ENABLED` - Enable CORS: "true" or "false" (default: "true")

## Usage Examples

### Using cURL

**Add two numbers:**
```bash
curl -X POST http://localhost:3000/add \
  -H "Content-Type: application/json" \
  -d '{"number1": 10, "number2": 5}'
```

**Subtract two numbers:**
```bash
curl -X POST http://localhost:3000/subtract \
  -H "Content-Type: application/json" \
  -d '{"number1": 10, "number2": 3}'
```

**Multiply two numbers:**
```bash
curl -X POST http://localhost:3000/multiply \
  -H "Content-Type: application/json" \
  -d '{"number1": 4, "number2": 6}'
```

**Divide two numbers:**
```bash
curl -X POST http://localhost:3000/divide \
  -H "Content-Type: application/json" \
  -d '{"dividend": 20, "divisor": 4}'
```

**Sum an array of numbers:**
```bash
curl -X POST http://localhost:3000/sum \
  -H "Content-Type: application/json" \
  -d '[1, 2, 3, 4, 5]'
```

## Project Structure

```
rest-api/
├── docs/
│   ├── api-spec.yaml    # OpenAPI 3.0 specification
│   └── README.md        # Project requirements and context
├── cmd/
│   └── api/
│       └── main.go      # Application entry point
├── internal/
│   ├── handlers/        # HTTP request handlers
│   ├── models/          # Request/response models
│   ├── services/        # Business logic
│   └── middleware/      # HTTP middleware (logging, CORS, validation)
├── pkg/
│   └── logger/          # Logging utilities
├── README.md            # This file
└── go.mod              # Go module definition
```

## Error Handling Examples

The API provides detailed error responses:

**Division by zero:**
```bash
curl -X POST http://localhost:3000/divide \
  -H "Content-Type: application/json" \
  -d '{"dividend": 10, "divisor": 0}'

# Response: 400 Bad Request
{
  "error": "Division by zero is not allowed",
  "code": "DIVISION_BY_ZERO"
}
```

**Invalid input:**
```bash
curl -X POST http://localhost:3000/add \
  -H "Content-Type: application/json" \
  -d '{"number1": "invalid", "number2": 5}'

# Response: 400 Bad Request
{
  "error": "Invalid input: number1 must be an integer",
  "code": "INVALID_INPUT"
}
```

## API Documentation

- **OpenAPI Spec**: See `docs/api-spec.yaml` for the complete API specification
- **Project Requirements**: See `docs/README.md` for detailed project context and requirements
- **Reference Implementation**: https://calculator.dreamsofcode.io

## Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o calculator-api ./cmd/api
./calculator-api
```

### Logging Examples

**Text Format (Development):**
```
2025-07-01T10:30:00Z INFO request completed method=POST path=/add status=200 ip=127.0.0.1 duration=2ms
```

**JSON Format (Production):**
```json
{"time":"2025-07-01T10:30:00Z","level":"INFO","msg":"request completed","method":"POST","path":"/add","status":200,"ip":"127.0.0.1","duration":"2ms"}
```

## Additional Features (Future Enhancements)

- **Rate Limiting**: Prevent API misuse with request rate limits
- **Authentication**: Token-based authentication for authorized access
- **Database Integration**: Store calculation history and analytics
- **Floating Point Support**: Extend operations to support decimal numbers
- **Client SDK**: HTTP client library for easy API integration
- **Web Frontend**: Interactive calculator interface
- **Request Tracing**: Add request IDs for distributed tracing

## Development

### Running Tests
```bash
go test ./...
```

### Building for Production
```bash
go build -o calculator-api
./calculator-api
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.